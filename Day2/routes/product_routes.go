package routes

import (
	"Day2/controllers"

	"github.com/gorilla/mux"
)

func SetupProductRoute() *mux.Router {

	routes := mux.NewRouter()

	routes.HandleFunc("/products", controllers.CreateProduct).Methods("POST") // ke wo apna product create kr ske

	routes.HandleFunc("/products", controllers.GetProducts).Methods("GET") // get all the products catalog

	return routes

}
