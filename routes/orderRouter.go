package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/pyakz/resto-app/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	orders := incomingRoutes.Group("/orders")

	orders.GET("/", controller.GetOrders())
	orders.GET("/:id", controller.GetOrder())
	orders.GET("/:id/order_items", controller.GetOrderItemsByOrderId())
	orders.POST("/", controller.CreateOrder())
	orders.PATCH("/:id", controller.UpdateOrder())
	orders.DELETE("/:id", controller.RemoveOrder())

}
