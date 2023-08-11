package routes

import (
	customerController "REST-API/Controller/Customer"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(r *gin.Engine) {
	r.GET("/api/v1/customers", customerController.GetAll)
	r.GET("/api/v1/customers/:id", customerController.GetACustomer)
	r.POST("/api/v1/customers", customerController.CreateCustomer)
	r.PATCH("/api/v1/customers/:id", customerController.UpdateCustomer)
	r.DELETE("/api/v1/customers/:id", customerController.DeleteCustomer)
}
