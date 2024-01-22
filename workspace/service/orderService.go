package service

import (
	"github.com/AdiLambe/TestGoLangCodes/workspace/domain"
	"github.com/AdiLambe/TestGoLangCodes/workspace/errs"
)

type OrderService interface {
	GetRepository() domain.OrderRepository
	CreateOrder(order domain.Order) (*domain.Order, *errs.AppError)
	// It will take a string and will return pointer to a domain order and error
	GetOrder(id string) (*domain.Order, *errs.AppError)
	SaveOrder(order domain.Order) (*domain.Order, *errs.AppError)
	GetOrdersList(status string) ([]domain.Order, *errs.AppError)
}

// service implementation
type DefaultOrderService struct {
	repo domain.OrderRepository
}

func (s DefaultOrderService) GetRepository() domain.OrderRepository {
	return s.repo
}

func (s DefaultOrderService) CreateOrder(order domain.Order) (*domain.Order, *errs.AppError) {
	return &order, nil
}

func (s DefaultOrderService) SaveOrder(order domain.Order) (*domain.Order, *errs.AppError) {
	return s.repo.SaveOrder(order)
}

func (s DefaultOrderService) GetOrdersList(status string) ([]domain.Order, *errs.AppError) {
	return s.repo.FindAll(status)
}

func (s DefaultOrderService) GetOrder(id string) (*domain.Order, *errs.AppError) {
	return s.repo.ById(id) //"s.repo.ById(id)" by making this call we have connected from primary port with the secondary port
}

func NewOrderService(repository domain.OrderRepository) OrderService {
	return DefaultOrderService{repository}
}
