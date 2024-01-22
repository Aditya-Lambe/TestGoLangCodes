package domain

import (
	"github.com/AdiLambe/TestGoLangCodes/workspace/errs"
)

type Order struct {
	Id          string `json:"order_id" gorm:"column:order_id"`
	Name        string `json:"order_name" gorm:"column:order_name"`
	Number      string `json:"order_number" gorm:"column:order_number"`
	Description string `json:"order_description" gorm:"column:order_description"`
	Status      string `json:"order_status" gorm:"column:status"`
}

// This is our interface
type OrderRepository interface {
	SaveOrder(order Order) (*Order, *errs.AppError)
	FindAll(status string) ([]Order, *errs.AppError)

	/* This method will take id of string type
	This "*Order" is a pointer to send nil in case there is no customer available against this id*/
	ById(id string) (*Order, *errs.AppError)
}
