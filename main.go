package main

import (
	"net/http"

	"github.com/FernandoDanielZigarra/rest-api-postgres/db"
	"github.com/FernandoDanielZigarra/rest-api-postgres/models"
	"github.com/FernandoDanielZigarra/rest-api-postgres/routes"
	"github.com/gorilla/mux"
)


func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Task{})

	r := mux.NewRouter()

	r.HandleFunc(("/"), routes.HomeHandler)
	r.HandleFunc(("/users"), routes.GetUsersHandler).Methods("GET")
	r.HandleFunc(("/users/{id}"), routes.GetUserHandler).Methods("GET")
	r.HandleFunc(("/users"), routes.PostUserHandler).Methods("POST")
	r.HandleFunc(("/users/{id}"), routes.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc(("/users/{id}"), routes.UpdateUserHandler).Methods("PUT")

	r.HandleFunc(("/tasks"), routes.GetTasksHandler).Methods("GET")
	r.HandleFunc(("/tasks/{id}"), routes.GetTaskHandler).Methods("GET")
	r.HandleFunc(("/tasks"), routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc(("/tasks/{id}"), routes.DeleteTaskHandler).Methods("DELETE")
	r.HandleFunc(("/tasks/{id}"), routes.UpdateTaskHandler).Methods("PUT")

	http.ListenAndServe(":3000", r)
}
