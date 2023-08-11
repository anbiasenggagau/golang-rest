package customerController

import (
	model "REST-API/Model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAll(context *gin.Context) {
	var customers []model.Customer

	model.Db.Find(&customers)

	context.IndentedJSON(http.StatusOK, customers)
}

func GetACustomer(context *gin.Context) {
	var customer model.Customer
	id := context.Param("id")

	if err := model.Db.First(&customer, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			context.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf(`Cannot find customer ID %s`, id)})
			return
		default:
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	context.IndentedJSON(http.StatusOK, customer)

}

func CreateCustomer(context *gin.Context) {
	var customer model.Customer

	if err := context.ShouldBindJSON(&customer); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	model.Db.Create(&customer)
	context.IndentedJSON(http.StatusCreated, customer)
}

func UpdateCustomer(context *gin.Context) {
	var customer model.Customer
	id := context.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	if err := context.ShouldBindJSON(&customer); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if model.Db.Model(&customer).Where("id = ?", id).Updates(&customer).RowsAffected == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Data is changed"})
}

func DeleteCustomer(context *gin.Context) {
	var customer model.Customer
	id := context.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	if model.Db.Delete(&customer, id).RowsAffected == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Data is deleted"})
}
