package main

import (
	"log"
	"net/http"
	"vacation-planner/database"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	//test database connection for default route
	//prints "error" on error, prints "successful connection" on success
	database.Connect()
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	//db := database.Connect()
	//h := routes.NewConnection(db)

	r := mux.NewRouter()

	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))
}
