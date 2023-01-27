package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"net/http"
    "github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

func main() {
	
	client, err := mongo.Connect(context.TODO(), clientOptions())
	
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	mongoDB := db.NewMongo(client)
	
	// CORS is enabled only in prod profile
	cors := os.Getenv("profile") == "prod"
	app := web.NewApp(mongoDB, cors)
	err = app.Serve()
	log.Println("Error", err)
}

func clientOptions() *options.ClientOptions {
	host := "db"
	if os.Getenv("profile") != "prod" {
		host = "localhost"
	}
	return options.Client().ApplyURI(
		"mongodb://" + host + ":27017",
	)
}
