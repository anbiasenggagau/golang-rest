package routes

import (
	customerAddressController "REST-API/Controller/CustomerAddress"

	"github.com/gin-gonic/gin"
)

func CustomerAddressRoutes(r *gin.Engine) {
	r.GET("/api/v1/address", customerAddressController.GetAll)
	r.GET("/api/v1/address/:id", customerAddressController.GetAddress)
	r.POST("/api/v1/address", customerAddressController.CreateAddress)
	r.PATCH("/api/v1/address/:id", customerAddressController.UpdateAddress)
	r.DELETE("/api/v1/address/:id", customerAddressController.DeleteAddress)
}
