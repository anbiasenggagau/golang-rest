package routes

import (
	productController "REST-API/Controller/Product"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/api/v1/products", productController.GetAll)
	r.GET("/api/v1/products/:id", productController.GetAProduct)
	r.POST("/api/v1/products", productController.CreateProduct)
	r.PATCH("/api/v1/products/:id", productController.UpdateProduct)
	r.DELETE("/api/v1/products/:id", productController.DeleteProduct)
}
