package model

import "time"

type Transaction struct {
	Id            int64     `gorm:"primaryKey" json:"id"`
	CustomerName  string    `json:"customer_name"`
	ProductName   string    `json:"product_name"`
	PaymentMethod string    `json:"payment_method"`
	TotalPrice    float32   `json:"total_price"`
	CreatedAt     time.Time `grom:"default:current_timestamp" json:"created_at"`
}
