package routes

import (
	transactionController "REST-API/Controller/Transaction"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(r *gin.Engine) {
	r.GET("/api/v1/transactions", transactionController.GetAll)
	r.GET("/api/v1/transactions/:id", transactionController.GetATransaction)
	r.POST("/api/v1/transactions", transactionController.CreateTransaction)
}
