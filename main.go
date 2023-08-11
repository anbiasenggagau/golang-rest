package main

import (
	model "REST-API/Model"
	routes "REST-API/Routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	model.ConnectDatabase()

	routes.CustomerRoutes(r)
	routes.CustomerAddressRoutes(r)
	routes.ProductRoutes(r)
	routes.PaymentMethodRoutes(r)
	routes.TransactionRoutes(r)

	r.Run("localhost:3030")
}
