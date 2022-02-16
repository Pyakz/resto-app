package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/pyakz/resto-app/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	users := incomingRoutes.Group("/users")
	users.POST("/", controller.CreateUser())
	users.GET("/", controller.GetUser())
	users.GET("/:id", controller.GetUser())
	users.POST("/signup", controller.LoginUser())
	users.POST("/login", controller.SignUpUser())
}
