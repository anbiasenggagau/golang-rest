package routes

import (
	paymentController "REST-API/Controller/PaymentMethod"

	"github.com/gin-gonic/gin"
)

func PaymentMethodRoutes(r *gin.Engine) {
	r.GET("/api/v1/payments", paymentController.GetAll)
	r.GET("/api/v1/payments/:id", paymentController.GetAPayment)
	r.POST("/api/v1/payments", paymentController.CreatePayment)
	r.PATCH("/api/v1/payments/:id", paymentController.UpdatePayment)
	r.DELETE("/api/v1/payments/:id", paymentController.DeletePayment)
}
