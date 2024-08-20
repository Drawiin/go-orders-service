package repository

import "github.com/drawiin/go-orders-service/internal/entity"

type OrderRepository interface {
	SaveOrder(order *entity.Order) error
	GetAllOrders() ([]*entity.Order, error)
	GetOrder(id string) (*entity.Order, error)
}
