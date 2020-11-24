package main

import (
	"os"
	"strconv"
	"time"

	"github.com/bmf-san/gobel-api/app/infrastructure"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	threshold, _ := strconv.Atoi(os.Getenv("LOG_THRESHOLD"))
	offset, _ := strconv.Atoi(os.Getenv("LOG_TIME_ZONE_OFFSET"))
	location := time.FixedZone(os.Getenv("TIME_ZONE"), offset)

	logger := infrastructure.NewLogger(threshold, location)
	db := infrastructure.NewDB()

	connMySQL, err := db.GetConnMySQL()
	if err != nil {
		logger.Error(err.Error())
	}

	connRedis, err := db.GetConnRedis()
	if err != nil {
		logger.Error(err.Error())
	}

	infrastructure.Dispatch(connMySQL, connRedis, logger)
}
