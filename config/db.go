package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var Client *mongo.Client


func GetDatabase() *mongo.Database {

	DB_NAME := os.Getenv("DB_NAME")

	return Client.Database(DB_NAME)

}

func ConnectDatabase() {

	MONGO_URI := os.Getenv("MONGO_URI")
	
	clientOption := options.Client().ApplyURI(MONGO_URI)

	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Println(err)
		log.Fatalf("Error while connect to MongoDB!")
	}

	Client = client

	log.Println("Successfully connected to MongoDB!")

}

func DisconnectDatabase() {

	err := Client.Disconnect(context.Background())
	if err != nil {
		log.Fatalf("Error while disconnect from MongoDB!")
	}

	log.Println("Successfully disconnected from MongoDB!")

}