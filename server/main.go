package main

import (
	"log"
	"net/http"
	"vacation-planner/database"
	"vacation-planner/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Establishing database connection
	db := database.Connect()

	// Packaging database connection to be passed to handlers
	h := routes.NewConnection(db)

	r := mux.NewRouter()

	// Created POST request for creating a user in the database
	r.HandleFunc("/createUser", h.CreateUser).Methods("POST")

	// Created GET request to get relevant travel information for the user
	r.HandleFunc("/newDestination/{location}", h.GetDestInfo).Methods("GET")

	// Created another GET to "get" a users password and ensure the password given matches our records
	r.HandleFunc("/loginUser", h.LoginUser).Methods("GET")

	// Created a new route that can take 3 HTTP method requests to update locations in the savedLocation table for users to have their own Destination Lists
	r.HandleFunc("/updateDestination", h.UpdateDestination).Methods("PUT")
	r.HandleFunc("/updateDestination", h.UpdateDestination).Methods("DELETE")
	r.HandleFunc("/updateDestination", h.UpdateDestination).Methods("GET")

	// Enabling CORS, binding to a port and passing our router in
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:4200"})
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(r)))
}
