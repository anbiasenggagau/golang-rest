package transactionController

import (
	model "REST-API/Model"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type request struct {
	CustomerName  string `json:"customer_name"`
	ProductName   string `json:"product_name"`
	PaymentMethod string `json:"payment_method"`
}

func GetAll(context *gin.Context) {
	var transactions []model.Transaction

	model.Db.Find(&transactions)

	context.IndentedJSON(http.StatusOK, transactions)
}

func GetATransaction(context *gin.Context) {
	var transaction model.Transaction
	id := context.Param("id")

	if err := model.Db.First(&transaction, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			context.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf(`Cannot find transaction ID %s`, id)})
			return
		default:
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	context.IndentedJSON(http.StatusOK, transaction)

}

func CreateTransaction(context *gin.Context) {
	var req request
	var transaction model.Transaction
	var productsData []model.Product
	var paymentsData []model.PaymentMethod
	var customer model.Customer

	var resProductName string
	var resPaymentMethod string
	var resTotalPrice float32

	if err := context.ShouldBindJSON(&req); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	products := strings.Split(req.ProductName, ",")
	payments := strings.Split(req.PaymentMethod, ",")

	model.Db.Where("(name) IN ?", products).Find(&productsData)
	if len(productsData) == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": `Cannot find specified product`})
		return
	}

	model.Db.Where("(name) IN ?", payments).Find(&paymentsData)
	if len(paymentsData) == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": `Cannot find specified payment method`})
		return
	}

	for _, item := range productsData {
		resTotalPrice += item.Price
		resProductName += item.Name + ","
	}

	for _, item := range paymentsData {
		if !item.IsActive {
			context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "one of the payment method is unactive"})
			return
		}
		resPaymentMethod += item.Name + ","
	}

	resProductName = resProductName[:len(resProductName)-1]
	resPaymentMethod = resPaymentMethod[:len(resPaymentMethod)-1]

	transaction.ProductName = resProductName
	transaction.PaymentMethod = resPaymentMethod
	transaction.TotalPrice = resTotalPrice

	model.Db.Create(&transaction)

	// Create new customer data if her/his name is not exist in customer table
	model.Db.Where("customer_name = ?", req.CustomerName).Find(&customer)
	if customer == (model.Customer{}) {
		var customer model.Customer
		customer.CustomerName = req.CustomerName
		model.Db.Create(&customer)

		transaction.CustomerName = customer.CustomerName
	} else {
		transaction.CustomerName = req.CustomerName
	}

	context.IndentedJSON(http.StatusOK, transaction)
}
