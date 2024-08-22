package graph

import "github.com/drawiin/go-orders-service/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase  *usecase.CreateOrderUseCase
	GetAllOrdersUseCase *usecase.GetAllOrdersUseCase
	GetOrderByIdUseCase *usecase.GetOrderByIdUseCase
}

func NewResolver(
	createOrderUseCase *usecase.CreateOrderUseCase,
	getAllOrdersUseCase *usecase.GetAllOrdersUseCase,
	getOrderByIdUseCase *usecase.GetOrderByIdUseCase,
) *Resolver {
	return &Resolver{
		CreateOrderUseCase:  createOrderUseCase,
		GetAllOrdersUseCase: getAllOrdersUseCase,
		GetOrderByIdUseCase: getOrderByIdUseCase,
	}
}
