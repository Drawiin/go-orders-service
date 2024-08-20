package usecase

import "github.com/drawiin/go-orders-service/internal/repository"

type GetOrderByIdUseCase struct {
	OrderRepository repository.OrderRepository
}

func NewGetOrderByIdUseCase(OrderRepository repository.OrderRepository) *GetOrderByIdUseCase {
	return &GetOrderByIdUseCase{
		OrderRepository: OrderRepository,
	}
}

func (u *GetOrderByIdUseCase) Execute(id string) (OrderOutputDTO, error) {
	order, err := u.OrderRepository.GetOrder(id)
	if err != nil {
		return OrderOutputDTO{}, err
	}

	return OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.Price,
	}, nil
}