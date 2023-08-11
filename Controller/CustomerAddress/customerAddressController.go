package customerAddressController

import (
	model "REST-API/Model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type addressResponse struct {
	Id         int64  `json:"id"`
	Address    string `json:"address"`
	CustomerId int64  `json:"customer_id"`
}

func GetAll(context *gin.Context) {
	var addresses []model.CustomerAddress

	model.Db.Find(&addresses)

	response := make([]addressResponse, 0)
	for _, item := range addresses {
		response = append(response, addressResponse{
			Id:         item.Id,
			CustomerId: item.CustomerId,
			Address:    item.Address,
		})
	}

	context.IndentedJSON(http.StatusOK, response)
}

func GetAddress(context *gin.Context) {
	var address model.CustomerAddress
	id := context.Param("id")

	if err := model.Db.First(&address, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			context.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf(`Cannot find customer address ID %s`, id)})
			return
		default:
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	var response addressResponse = addressResponse{
		Id:         address.Id,
		CustomerId: address.CustomerId,
		Address:    address.Address,
	}

	context.IndentedJSON(http.StatusOK, response)

}

func CreateAddress(context *gin.Context) {
	var address model.CustomerAddress

	if err := context.ShouldBindJSON(&address); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := model.Db.Create(&address).Error
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var response addressResponse = addressResponse{
		Id:         address.Id,
		CustomerId: address.CustomerId,
		Address:    address.Address,
	}

	context.IndentedJSON(http.StatusCreated, response)
}

func UpdateAddress(context *gin.Context) {
	var address model.CustomerAddress
	id := context.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	if err := context.ShouldBindJSON(&address); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := model.Db.Model(&address).Where("id = ?", id).Updates(&address).Error; err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Data is changed"})
}

func DeleteAddress(context *gin.Context) {
	var address model.CustomerAddress
	id := context.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	if model.Db.Delete(&address, id).RowsAffected == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Data is deleted"})
}
