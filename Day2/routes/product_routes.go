package routes

import (
	"Day2/controllers"
	"Day2/middleware"

	"github.com/gorilla/mux"
)

func SetupProductRoute() *mux.Router {

	routes := mux.NewRouter()

	routes.HandleFunc("/login", controllers.Login).Methods("POST") // public route for login

	api := routes.PathPrefix("/api").Subrouter() // mean all other routes and endpoints now protected by this jwt token required for access that credentials

	api.Use(middleware.ValidateJWT) // now use jwt to validate the login then access it to protected routes which are given below

	api.HandleFunc("/products", controllers.CreateProduct).Methods("POST") // ke wo apna product create kr ske

	api.HandleFunc("/products", controllers.GetProducts).Methods("GET") // get all the products catalog

	api.HandleFunc("/products/{id}", controllers.GetProduct).Methods("GET") // get all the products catalog

	api.HandleFunc("/products/{id}", controllers.UpdateTask).Methods("PUT") // get all the products catalog

	api.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE") // get all the products catalog

	return routes

}
