package models

import "github.com/jinzhu/gorm"

//Order ...
type Order struct {
	gorm.Model
	OrderID       string `gorm:"type:varchar(36);unique_index"`
	TransactionID string
	//TODO wait for more info
	Currency string
}
