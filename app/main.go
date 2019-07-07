package main

import (
	"os"
	"time"

	"github.com/bmf-san/gobel-api/app/infrastructure"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	location := os.Getenv("TIME_ZONE")
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func main() {
	logger := infrastructure.NewLogger()
	db := infrastructure.NewDB()

	conn, err := db.GetConn()
	if err != nil {
		logger.LogError(err)
	}

	infrastructure.Dispatch(conn, logger)
}
