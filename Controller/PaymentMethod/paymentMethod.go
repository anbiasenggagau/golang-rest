package paymentController

import (
	model "REST-API/Model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAll(context *gin.Context) {
	var payments []model.PaymentMethod

	model.Db.Find(&payments)

	context.IndentedJSON(http.StatusOK, payments)
}

func GetAPayment(context *gin.Context) {
	var payment model.PaymentMethod
	id := context.Param("id")

	if err := model.Db.First(&payment, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			context.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf(`Cannot find payment ID %s`, id)})
			return
		default:
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	context.IndentedJSON(http.StatusOK, payment)

}

func CreatePayment(context *gin.Context) {
	var payment model.PaymentMethod

	if err := context.ShouldBindJSON(&payment); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	model.Db.Create(&payment)
	context.IndentedJSON(http.StatusCreated, payment)
}

func UpdatePayment(context *gin.Context) {
	var payment model.PaymentMethod
	id := context.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	if err := context.ShouldBindJSON(&payment); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if model.Db.Model(&payment).Where("id = ?", id).Updates(&payment).RowsAffected == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Data is changed"})
}

func DeletePayment(context *gin.Context) {
	var payment model.PaymentMethod
	id := context.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	if model.Db.Delete(&payment, id).RowsAffected == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Data is deleted"})
}
