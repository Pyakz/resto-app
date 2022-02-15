package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pyakz/resto-app/middleware"
	"github.com/pyakz/resto-app/routes"
)

func main() {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + PORT)

}
