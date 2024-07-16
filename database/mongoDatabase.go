package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	uri := os.Getenv("DB_URL")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal("Failed to connect to Mongo database! \n", err.Error())
		os.Exit(2)
	} else {
		log.Println("Connected to Mongo database successfully!")
	}

	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	var result bson.M
	if err := client.Database("aasan").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return client
}

var Client *mongo.Client = ConnectToMongo()

func OpenConnection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("aasan").Collection(collectionName)
	return collection
}
