package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/pyakz/resto-app/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	foods := incomingRoutes.Group("/foods")

	foods.GET("/", controller.GetFoods())
	foods.GET("/:id", controller.GetFood())
	foods.POST("/", controller.CreateFood())
	foods.PATCH("/:id", controller.UpdateFood())
	foods.DELETE("/:id", controller.RemoveFood())

}
