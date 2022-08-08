package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"strconv"
	"syscall"
	"time"

	"github.com/bmf-san/gobel-api/app/infrastructure"
	"github.com/bmf-san/gobel-api/app/infrastructure/database"
	_ "github.com/go-sql-driver/mysql"
)

const timeout time.Duration = 10 * time.Second

func main() {
	threshold, _ := strconv.Atoi(os.Getenv("LOG_THRESHOLD"))
	offset, _ := strconv.Atoi(os.Getenv("LOG_TIME_ZONE_OFFSET"))
	location := time.FixedZone(os.Getenv("TIME_ZONE"), offset)

	logger := infrastructure.NewLogger(threshold, location)

	defer func() {
		if x := recover(); x != nil {
			logger.Error(string(debug.Stack()))
		}
		os.Exit(1)
	}()

	mc := database.NewMySQLConn()
	connm, err := mc.Conn()
	if err != nil {
		logger.Error(err.Error())
	}

	rc := database.NewRedisConn()
	connr, err := rc.Conn()
	if err != nil {
		logger.Error(err.Error())
	}

	r := infrastructure.Route(connm, connr, logger)

	s := http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: r,
	}

	go func() {
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			logger.Error(err.Error())
		}
	}()

	q := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-q

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Gracefully shutdown")
}
