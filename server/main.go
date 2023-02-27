package main

import (
	"log"
	"net/http"
	"vacation-planner/database"
	"vacation-planner/routes"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func main() {
	// Establishing database connection
	db := database.Connect()

	// Packaging database connection to be passed to handlers
	h := routes.NewConnection(db)

	r := mux.NewRouter()

	// Created test POST request for database
	r.HandleFunc("/createUser", h.CreateUser).Methods("POST")

	r.HandleFunc("/newDestination/{location}", h.GetDestInfo).Methods("GET")

	r.HandleFunc("/loginUser", h.LoginUser).Methods("GET")

	// Enabling CORS, binding to a port and passing our router in
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
    origins := handlers.AllowedOrigins([]string{"http://localhost:4200"})
    log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(r)))
}
