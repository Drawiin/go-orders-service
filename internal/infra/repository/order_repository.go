package repository

import (
	"context"

	"github.com/drawiin/go-orders-service/internal/entity"
	"github.com/drawiin/go-orders-service/internal/infra/db"
)

type OrderRepository struct {
	Db *db.Queries
}

func NewOrderRepository(db *db.Queries) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) SaveOrder(order *entity.Order) error {
	err := r.Db.CreateOrder(context.Background(), db.CreateOrderParams{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetAllOrders() ([]*entity.Order, error) {
	orders, err := r.Db.ListOrders(context.Background())
	if err != nil {
		return nil, err
	}

	var ordersList []*entity.Order
	for _, order := range orders {
		ordersList = append(ordersList, &entity.Order{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return ordersList, nil
}

func (r *OrderRepository) GetOrder(id string) (*entity.Order, error) {
	return nil, nil
}
