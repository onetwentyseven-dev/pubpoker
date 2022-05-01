package main

import (
	"database/sql"
	"log"
	"time"

	driver "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func loadMySQL() {

	mysqlConfig := driver.Config{
		User:                 ssmConfig.MySQLUsername,
		Passwd:               ssmConfig.MySQLPassword,
		Net:                  "tcp",
		Addr:                 ssmConfig.MySQLHost,
		DBName:               ssmConfig.MySQLDB,
		Loc:                  time.UTC,
		Timeout:              time.Second * 5,
		ReadTimeout:          time.Second * 5,
		WriteTimeout:         time.Second * 5,
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		log.Panicf("[MySQL Connect] Failed to connect to mysql server: %s", err)
	}

	db.SetConnMaxIdleTime(time.Second * 5)
	db.SetConnMaxLifetime(time.Second * 30)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(15)

	for i := 0; i < 3; i++ {
		err = db.Ping()
		if err != nil {
			logger.WithError(err).Error("[MySQL Connect] Failed to ping mysql server")
		}

		if err == nil {
			break
		}

	}
	dbConn = sqlx.NewDb(db, "mysql")

}
