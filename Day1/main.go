package main

import (
	"Day1/database"
	"Day1/routes"
	"log"
	"net/http"
)

func main() {

	database.Connect()
	log.Println("=====DB Connected Successfully====") // connecting with db

	r := routes.SetupRoute() // now call the routes

	log.Println("Server Started on Port : 8080")
	log.Fatal(http.ListenAndServe(":8080", r)) // running and serve on port 8080 lets check it out

}
