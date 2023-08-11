package model

type PaymentMethod struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}
