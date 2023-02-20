package main

import (
	"log"
	"net/http"
	"vacation-planner/database"
	"vacation-planner/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Establishing database connection
	db := database.Connect()

	// Packaging database connection to be passed to handlers
	h := routes.NewConnection(db)

	r := mux.NewRouter()

	// Created test POST request for database
	r.HandleFunc("/CreateUser", h.CreateUser).Methods("POST")

	r.HandleFunc("/newDestination/{location}", h.GetDestInfo).Methods("GET")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))
}
