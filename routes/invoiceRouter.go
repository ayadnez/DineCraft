package routes

import (
	controller "PROJECT_1/controller"

	"github.com/gin-gonic/gin"
)

func IvoiceRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controller.GetInvoice())
	incomingRoutes.POST("/invoices", controller.CreateInvoice())
	incomingRoutes.PATCH("invoices/:invoice_id", controller.UpdateInvoice())
}
