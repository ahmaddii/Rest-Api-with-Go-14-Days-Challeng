package routes

import (
	"Day1/controllers"
	"Day1/middleware"

	"github.com/gorilla/mux"
)

func SetupRoute() *mux.Router {

	routes := mux.NewRouter() // create routes for the endpoint of api

	routes.HandleFunc("/login", controllers.Login).Methods("POST") // Public Route anyone can login and can get the jwt token 

	api := routes.PathPrefix("/api").Subrouter() // all other routes required the jwt token and all are protected


	// you can use subrouter inside the routes but the upper one is resusable by pathprefix

	// which is like api/tasks now become this url

	api.Use(middleware.ValidateJWT) // use jwt for validation // with this line all routes are protected
	

	// now all the routes enpoint becomes like api/tasks , api/tasks{id} ...

	api.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")        // post and create tasks
	api.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")           // Get all the tasks
	api.HandleFunc("/tasks/{id}", controllers.GetTaskByID).Methods("GET")   // Get Specifc Id task
	api.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")    // update specifc task
	api.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE") // delete a task

	return routes // return all routes
}
