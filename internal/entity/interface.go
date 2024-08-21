package entity

type OrderRepositoryInterface interface {
	SaveOrder(order *Order) error
	GetAllOrders() ([]*Order, error)
	GetOrder(id string) (*Order, error)
}
