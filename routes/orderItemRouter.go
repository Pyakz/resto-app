package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/pyakz/resto-app/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	order_items := incomingRoutes.Group("/order_items")
	order_items.GET("/", controller.GetOrderItems())
	order_items.GET("/:id", controller.GetOrderItem())
	order_items.POST("/", controller.CreateOrderItem())
	order_items.PATCH("/:id", controller.UpdateOrderItem())
	order_items.DELETE("/:id", controller.RemoveOrderItem())
}
