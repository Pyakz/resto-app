package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/pyakz/resto-app/controllers"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	invoice := incomingRoutes.Group("/invoice")

	invoice.GET("/", controller.GetInvoices())
	invoice.GET("/:id", controller.GetInvoice())
	invoice.POST("/", controller.CreateInvoice())
	invoice.PATCH("/:id", controller.UpdateInvoice())
	invoice.DELETE("/:id", controller.RemoveInvoice())

}
