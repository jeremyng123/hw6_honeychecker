/*
Package database for all database matters
*/
package database

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

/*
Client : Persistence mongo connection
*/
var Client *mongo.Client; 
var err error;

func init() {
	Client, err = GetMongoDBConnection();
}

//GetMongoDBConnection get connection of mongodb
// https://dev.to/joojodontoh/build-an-api-endpoint-with-golang-gin-2065 
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

	MongoDB := os.Getenv("MONGODB_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))
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

func GetMongoDBCollection(Client *mongo.Client, CollectionName string) (*mongo.Collection, error) {
	DBName := os.Getenv("DB_NAME")
	collection := Client.Database(DBName).Collection(CollectionName)
	return collection, nil
}


