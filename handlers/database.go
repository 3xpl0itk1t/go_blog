package handlers

import (
	"context"
	"fmt"
	"go_blog/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var user_collection *mongo.Collection
var blog_collection *mongo.Collection
var client *mongo.Client

func ConnectToDB() {
	//connection to the database

	clientOptions := options.Client().ApplyURI(config.MONGO_URL)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect To MongoDB Atlas
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Ping MongoDB Atlas
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB Atlas!")
	user_collection = client.Database("Blog").Collection("users")
	blog_collection = client.Database("Blog").Collection("blogs")

}

func DisconnectFromDB() {
	// Disconnect from MongoDB Atlas
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Disconnected from MongoDB Atlas!")
}
