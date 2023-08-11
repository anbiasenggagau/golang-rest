package model

type Product struct {
	Id    int64   `gorm:"primaryKey" json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
