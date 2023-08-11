package model

type CustomerAddress struct {
	Id         int64    `gorm:"primaryKey" json:"id"`
	Address    string   `json:"address"`
	CustomerId int64    `gorm:"uniqueIndex:customer_id" json:"customer_id"`
	Customer   Customer `gorm:"foreignKey:CustomerId;refrence:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
