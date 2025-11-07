package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {

	dsn := os.Getenv("MYSQL_DSN")

	// if not found connection string
	// then do this

	if dsn == "" {

		dsn = "backend-programmer:example-password@tcp(host.docker.internal:3306)/product"

	}

	var err error

	DB, err = sql.Open("mysql", dsn) // now open the database with current dsn connection

	if err != nil {

		log.Fatalf("Erros While Opening DB %v", err)

	}

	// check also during testing connection if error occur or not

	if err = DB.Ping(); err != nil {

		log.Fatalf("Error while testing the DB Connection %v", err)
	}

	fmt.Println("=====Successfully Connected to Mysql=====")

}
