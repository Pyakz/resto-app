package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/pyakz/resto-app/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		fmt.Println("----------------------GET FOODS-----------------------")
		fmt.Println(ctx, cancel)
		c.String(http.StatusOK, "Hello %s", "World")
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		fmt.Printf("----------------------GET FOOD: %v-----------------------", id)

		c.String(http.StatusOK, "Food ID: %s", id)
	}
}

func RemoveFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		fmt.Printf("----------------------REMOVE FOOD: %v-----------------------", id)

		c.String(http.StatusOK, "Food ID: %s", id)

	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		fmt.Printf("----------------------UPDATE FOOD: %v-----------------------", id)

		c.String(http.StatusOK, "Food ID: %s", id)

	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// func round(num float64) int {

// }

// func toFixed(num int64, precision int) int64 {

// }
