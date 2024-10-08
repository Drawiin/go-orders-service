package usecase

import "github.com/drawiin/go-orders-service/internal/entity"

type GetAllOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetAllOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *GetAllOrdersUseCase {
	return &GetAllOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (u *GetAllOrdersUseCase) Execute() (AllOrdersOutputDTO, error) {
	orders, err := u.OrderRepository.GetAllOrders()
	if err != nil {
		return AllOrdersOutputDTO{}, err
	}

	var ordersDTO AllOrdersOutputDTO
	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		}
		ordersDTO.Orders = append(ordersDTO.Orders, dto)
	}

	return ordersDTO, nil
}
