//go:build wireinject
// +build wireinject

package main

import (
	"github.com/drawiin/go-orders-service/cmd/ordersystem/di"
	"github.com/drawiin/go-orders-service/internal/infra/db"
	"github.com/drawiin/go-orders-service/internal/infra/graph"
	"github.com/drawiin/go-orders-service/internal/infra/web/web_handler"
	"github.com/drawiin/go-orders-service/pkg/events"
	"github.com/google/wire"
)

func NewWebOrderHandler(db db.DBTX, eventDispatcher events.EventDispatcherInterface) *web_handler.WebOrderHandler {
	wire.Build(
		di.WebOrderHandlerSet,
	)
	return &web_handler.WebOrderHandler{}
}

func NewGraphQLResolver(db db.DBTX, eventDispatcher events.EventDispatcherInterface) *graph.Resolver {
	wire.Build(
		di.GraphQLResolverSet,
	)
	return &graph.Resolver{}
}
