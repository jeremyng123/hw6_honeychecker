package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id"`
	UserID string `json:"user_id"`
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password"`
	Password1 string `json:"password1,omitempty"`
	Password2 string `json:"password2,omitempty"`
	Password3 string `json:"password3,omitempty"`
	Password4 string `json:"password4,omitempty"`
	Password5 string `json:"password5,omitempty"`
	Password6 string `json:"password6,omitempty"`
	Password7 string `json:"password7,omitempty"`
	Password8 string `json:"password8,omitempty"`
	Password9 string `json:"password9,omitempty"`
	Password10 string `json:"password10,omitempty"`
	Password11 string `json:"password11,omitempty"`
	Password12 string `json:"password12,omitempty"`
	Password13 string `json:"password13,omitempty"`
	Password14 string `json:"password14,omitempty"`
	Password15 string `json:"password15,omitempty"`
	Password16 string `json:"password16,omitempty"`
	Password17 string `json:"password17,omitempty"`
	Password18 string `json:"password18,omitempty"`
	Password19 string `json:"password19,omitempty"`
	CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}


func Register(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)

	var user User

	// bind request to data struct
	if err := c.BindJSON(&user); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		cancel()
		return
	}

	// validate the value fits the struct
	validationErr := validator.New().Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		cancel()
		return
	}

	// parse the timestamps on creation
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	user.ID = primitive.NewObjectID()
	user.UserID = user.ID.Hex()

	// insert the struct into collection
	result, err := UserCollection.InsertOne(ctx, user)
	if err != nil {
		msg := fmt.Sprintf("User %v is not created", user.Username)
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		cancel()
		return
	}
	defer cancel()

	// return ID of the created object to the frontend
	c.JSON(http.StatusOK, result)
}