package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Db *mongo.Database

func ConnectDb() {

	Client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := Client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("MongoDB ping failed: ", err)
	}
	Db = Client.Database("hakistream")
	log.Println("connected")
}
