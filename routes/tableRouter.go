package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/pyakz/resto-app/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	tables := incomingRoutes.Group("/tables")

	tables.GET("/", controller.GetTables())
	tables.GET("/:id", controller.GetTable())
	tables.POST("/", controller.CreateTable())
	tables.PATCH("/:id", controller.UpdateTable())
	tables.DELETE("/:id", controller.RemoveTable())
}
