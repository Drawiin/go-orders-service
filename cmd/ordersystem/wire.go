package main

import (
	"github.com/drawiin/go-orders-service/internal/entity"
	"github.com/drawiin/go-orders-service/internal/event"
	"github.com/drawiin/go-orders-service/internal/infra/db"
	"github.com/drawiin/go-orders-service/internal/infra/repository"
	"github.com/drawiin/go-orders-service/internal/usecase"
	"github.com/drawiin/go-orders-service/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	db.New,
	repository.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*repository.OrderRepository)),
)

var EventDispatcherSet = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

func NewCreateOrderUseCase(db db.DBTX) *usecase.CreateOrderUseCase {
	wire.Build(setOrderRepositoryDependency, EventDispatcherSet, usecase.NewCreateOrderUseCase)
	return &usecase.CreateOrderUseCase{}
}
