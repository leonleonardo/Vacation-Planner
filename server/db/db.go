package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"vacation-planner/models"
)

type DB interface {
	Destinations() ([]*model.Destination, error)
}

type MongoDB struct {
	collection *mongo.Collection
}

func NewMongo(client *mongo.Client) DB {
	dest := client.Database("dest").Collection("dest")
	return MongoDB{collection: dest}
}

func (m MongoDB) Destinations() ([]*model.Destination, error) {
	res, err := m.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Error while fetching destinations:", err.Error())
		return nil, err
	}
	var tech []*model.Technology
	err = res.All(context.TODO(), &tech)
	if err != nil {
		log.Println("Error while decoding destinations:", err.Error())
		return nil, err
	}
	return tech, nil
}
