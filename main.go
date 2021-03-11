package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jeremyng123/hw6_honeychecker/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection
var collectionErr error

// init get all collections for all package main to access
func init() {
	UserCollection, collectionErr = database.GetMongoDBCollection(database.Client, "user")
	if collectionErr != nil {
		log.Fatal(collectionErr)
	}
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", func (c *gin.Context){
		c.JSON(http.StatusOK, "Hello")
	})
	router.POST("/register", Register)
	

	router.Run(":" + port)
}