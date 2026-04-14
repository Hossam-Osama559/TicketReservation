package db

import (
	"TicketRservation/env"
	"crypto/tls"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dsn        = ""
	DbInstance *sql.DB
)

func InitDb() {

	mysql.RegisterTLSConfig("skip", &tls.Config{

		InsecureSkipVerify: true,
	})

	usr := env.DbUser.GetValue()
	password := env.DbPassword.GetValue()
	host := env.DbHost.GetValue()
	dbname := env.DbName.GetValue()
	port := env.DbPort.GetValue()

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&tls=skip", usr, password, host, port, dbname)

	var err error = nil

	DbInstance, err = sql.Open("mysql", dsn)

	if err != nil {

		panic(err)
	}
}
