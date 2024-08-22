//go:build wireinject
// +build wireinject

package main

import (
	"github.com/drawiin/go-orders-service/cmd/ordersystem/di"
	"github.com/drawiin/go-orders-service/internal/infra/db"
	"github.com/drawiin/go-orders-service/internal/infra/web/web_handler"
	"github.com/drawiin/go-orders-service/pkg/events"
	"github.com/google/wire"
)


func NewWebOrderHandler(db db.DBTX, eventDispatcher events.EventDispatcherInterface) *web_handler.WebOrderHandler {
	wire.Build(
		di.OrderRepositorySet,
		di.CreateOrderUseCaseSet,
		di.GetAllOrdersUseCaseSet,
		di.GetOrderByIdUseCaseSet,
		web_handler.NewWebOrderHandler,
	)
	return &web_handler.WebOrderHandler{}
}
