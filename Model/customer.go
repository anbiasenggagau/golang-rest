package model

type Customer struct {
	Id           int64  `gorm:"primaryKey" json:"id"`
	CustomerName string `json:"customer_name"`
}
