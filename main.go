package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pyakz/resto-app/database"
	"github.com/pyakz/resto-app/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	// router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + PORT)

}
