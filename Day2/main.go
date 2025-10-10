package main

import (
	"Day2/database"
	"Day2/routes"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	database.Connect() // call database func

	// print to log for debuging
	log.Println("====Database Connected Succesfully====")

	r := routes.SetupProductRoute() // call setup route fun

	log.Println("Server Started on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}
