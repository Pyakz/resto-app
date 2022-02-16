package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/pyakz/resto-app/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	menus := incomingRoutes.Group("/menus")
	
	menus.GET("/", controller.GetMenus())
	menus.GET("/:id", controller.GetMenu())
	menus.POST("/", controller.CreateMenu())
	menus.PATCH("/:id", controller.UpdateMenu())
	menus.DELETE("/:id", controller.RemoveMenu())

}
