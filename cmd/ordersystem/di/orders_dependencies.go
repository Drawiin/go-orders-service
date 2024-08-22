package di

import (
	"github.com/drawiin/go-orders-service/internal/entity"
	"github.com/drawiin/go-orders-service/internal/event"
	"github.com/drawiin/go-orders-service/internal/infra/db"
	"github.com/drawiin/go-orders-service/internal/infra/graph"
	grpc_service "github.com/drawiin/go-orders-service/internal/infra/grpc/service"
	"github.com/drawiin/go-orders-service/internal/infra/repository"
	"github.com/drawiin/go-orders-service/internal/infra/web/web_handler"
	"github.com/drawiin/go-orders-service/internal/usecase"
	"github.com/drawiin/go-orders-service/pkg/events"
	"github.com/google/wire"
)

var OrderRepositorySet = wire.NewSet(
	db.New,
	repository.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*repository.OrderRepository)),
)

var EventDispatcherSet = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

var CreateOrderUseCaseSet = wire.NewSet(
	EventDispatcherSet,
	usecase.NewCreateOrderUseCase,
)

var GetAllOrdersUseCaseSet = wire.NewSet(
	usecase.NewGetAllOrdersUseCase,
)

var GetOrderByIdUseCaseSet = wire.NewSet(
	usecase.NewGetOrderByIdUseCase,
)

var WebOrderHandlerSet = wire.NewSet(
	OrderRepositorySet,
	CreateOrderUseCaseSet,
	GetAllOrdersUseCaseSet,
	GetOrderByIdUseCaseSet,
	web_handler.NewWebOrderHandler,
)

var GraphQLResolverSet = wire.NewSet(
	OrderRepositorySet,
	CreateOrderUseCaseSet,
	GetAllOrdersUseCaseSet,
	GetOrderByIdUseCaseSet,
	graph.NewResolver,
)

var GrpcServiceSet = wire.NewSet(
	OrderRepositorySet,
	CreateOrderUseCaseSet,
	GetAllOrdersUseCaseSet,
	GetOrderByIdUseCaseSet,
	grpc_service.NewOrderService,
)
