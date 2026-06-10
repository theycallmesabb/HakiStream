package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

func ConnectDb() {

	Client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := Client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("MongoDB ping failed: ", err)
	}
	Db = Client.Database("hakistream")
	log.Println("connected")
}
