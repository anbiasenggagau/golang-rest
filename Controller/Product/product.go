package productController

import (
	model "REST-API/Model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAll(context *gin.Context) {
	var products []model.Product

	model.Db.Find(&products)

	context.IndentedJSON(http.StatusOK, products)
}

func GetAProduct(context *gin.Context) {
	var product model.Product
	id := context.Param("id")

	if err := model.Db.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			context.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf(`Cannot find product ID %s`, id)})
			return
		default:
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	context.IndentedJSON(http.StatusOK, product)

}

func CreateProduct(context *gin.Context) {
	var product model.Product

	if err := context.ShouldBindJSON(&product); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	model.Db.Create(&product)
	context.IndentedJSON(http.StatusCreated, product)
}

func UpdateProduct(context *gin.Context) {
	var product model.Product
	id := context.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	if err := context.ShouldBindJSON(&product); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if model.Db.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Data is changed"})
}

func DeleteProduct(context *gin.Context) {
	var product model.Product
	id := context.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	if model.Db.Delete(&product, id).RowsAffected == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Data is deleted"})
}
