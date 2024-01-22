package domain

import "github.com/AdiLambe/TestGoLangCodes/workspace/errs"

type OrderRepositoryStub struct {
	orders []Order
}

func (s OrderRepositoryStub) FindAll(status string) ([]Order, *errs.AppError) {
	return s.orders, nil
}

func NewOrderRepositoryStub() OrderRepositoryStub {
	orders := []Order{
		{"1001", "Nokia", "CPS3562X", "Mobile Phone", "1"},
		{"1022", "Lenovo", "SFG1562", "Laptop", "0"},
	}
	return OrderRepositoryStub{orders: orders}
}
