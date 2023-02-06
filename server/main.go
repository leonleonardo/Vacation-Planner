package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
    //"encoding/json"
    //"github.com/jmatth11/yfusion"
    //"fmt"
    //"os"
    //"github.com/joho/godotenv"
    //"gorm.io/gorm"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla!\n"))
}

func main() {

    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", YourHandler)

    // New function handler for "GET" requests
    r.HandleFunc("/newDestination/{location}", GetDestInfo).Methods("GET")

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8080", r))
}