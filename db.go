package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetMongoDBConnection get connection of mongodb
/*
Run the following command first in order to be able to connect to mongodb locally!
docker run --name mongo-db -p 27017:27017 -d mongo:latest
*/
func GetMongoDBConnection() (*mongo.Client, error) {
	err:=godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	MongoDb := os.Getenv("MONGODB_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
			log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println("Connected to MongoDB!")

	return client, nil
}

func GetMongoDBCollection(DBName string, CollectionName string) (*mongo.Collection, error) {
	client, err := GetMongoDBConnection()

	if err != nil {
		return nil, err
	}

	collection := client.Database(DBName).Collection(CollectionName)

	return collection, nil
}
