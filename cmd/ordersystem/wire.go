//go:build wireinject
// +build wireinject

package main

import (
	"github.com/drawiin/go-orders-service/internal/entity"
	"github.com/drawiin/go-orders-service/internal/event"
	"github.com/drawiin/go-orders-service/internal/infra/db"
	"github.com/drawiin/go-orders-service/internal/infra/repository"
	"github.com/drawiin/go-orders-service/internal/infra/web"
	"github.com/drawiin/go-orders-service/internal/usecase"
	"github.com/drawiin/go-orders-service/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	repository.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*repository.OrderRepository)),
)

var setEventsDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *db.Queries, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewGetAllOrdersUseCase(db *db.Queries) *usecase.GetAllOrdersUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewGetAllOrdersUseCase,
	)
	return &usecase.GetAllOrdersUseCase{}
}

func NewGetOrderByIdUseCase(db *db.Queries) *usecase.GetOrderByIdUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewGetOrderByIdUseCase,
	)
	return &usecase.GetOrderByIdUseCase{}
}

func NewWebOrderHandler(db *db.Queries, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
