package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoDB := os.Getenv("MONGODB_URI")

	if MongoDB == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}

	fmt.Println("MongoDB URI:", MongoDB)

	clientOptions := options.Client().ApplyURI(MongoDB)

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	return client

}

var Client *mongo.Client = DBinstance()

func OpenCollection(collectionName string) *mongo.Collection {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		log.Fatal("You must set your 'DATABASE_NAME' environmental variable.")
	}
	fmt.Println("Database Name:", databaseName)

	collection := Client.Database(databaseName).Collection(collectionName)

	if collection == nil {
		log.Fatal("Failed to open collection")
	}

	return collection
}
