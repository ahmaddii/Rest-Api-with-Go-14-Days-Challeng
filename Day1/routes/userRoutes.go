package routes

import (
	"Day1/controllers"

	"github.com/gorilla/mux"
)

func SetupRoute() *mux.Router {

	routes := mux.NewRouter() // create routes for the endpoint of api

	routes.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")      // post and create tasks
	routes.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")         // Get all the tasks
	routes.HandleFunc("/tasks/{id}", controllers.GetTaskByID).Methods("GET") // Get Specifc Id task

	routes.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT") // update specifc task

	routes.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE") // delete a task

	return routes // return all routes
}
