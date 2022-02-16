package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/pyakz/resto-app/controllers"
)

func NoteRoutes(incomingRoutes *gin.Engine) {
	notes := incomingRoutes.Group("/notes")

	notes.GET("/", controller.GetNotes())
	notes.GET("/:id", controller.GetNote())
	notes.POST("/", controller.CreateNote())
	notes.PATCH("/:id", controller.UpdateNote())
	notes.DELETE("/:id", controller.RemoveNote())

}
