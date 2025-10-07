package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)


var DB*sql.DB



func connect()  {

	dsn := os.Getenv("MYSQL_DSN") // enviroment var ke ander ja kr connection string dekhe ge


	 // if not found use our default dsn

    if dsn == "" {

		dsn = "backend-programmer:example-password@tcp(127.0.0.1:3306)/todo"
		
	}

	var err error

	DB,err = sql.Open("mysql",dsn) // open new database connection here

	// now we hanle errors

	if err != nil {

		log.Fatalf("Error Occurs during DB Connection %v",err) // it acutually prepares connection

	}

	// both are connection level error handling

	if err = DB.Ping(); err != nil { // check wheather error is not happend during testing
 
		log.Fatalf("Error while Testing %v",err)

	}

	fmt.Println("===== Connect to Mysql Successfully =====")
	
}