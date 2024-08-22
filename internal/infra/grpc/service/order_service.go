package grpc_service

import (
	"context"
	"io"

	"github.com/drawiin/go-orders-service/internal/infra/grpc/pb"
	"github.com/drawiin/go-orders-service/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase  *usecase.CreateOrderUseCase
	GetAllOrdersUseCase *usecase.GetAllOrdersUseCase
	GetOrderByIdUseCase *usecase.GetOrderByIdUseCase
}

func NewOrderService(
	createOrderUseCase *usecase.CreateOrderUseCase,
	getAllOrdersUseCase *usecase.GetAllOrdersUseCase,
	getOrderByIdUseCase *usecase.GetOrderByIdUseCase,
) *OrderService {
	return &OrderService{
		CreateOrderUseCase:  createOrderUseCase,
		GetAllOrdersUseCase: getAllOrdersUseCase,
		GetOrderByIdUseCase: getOrderByIdUseCase,
	}
}

func (s *OrderService) CreateOrder(context context.Context, request *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    request.Id,
		Price: float64(request.Price),
		Tax:   float64(request.Tax),
	}
	order, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.OrderResponse{
		Id:         order.ID,
		Price:      float32(order.Price),
		Tax:        float32(order.Tax),
		FinalPrice: float32(order.FinalPrice),
	}, nil
}
func (s *OrderService) CreateOrderStreamBidirectional(stream pb.OrderService_CreateOrderStreamBidirectionalServer) error {
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		dto := usecase.OrderInputDTO{
			ID:    order.Id,
			Price: float64(order.Price),
			Tax:   float64(order.Tax),
		}
		createdOrder, err := s.CreateOrderUseCase.Execute(dto)
		if err != nil {
			return err
		}
		err = stream.Send(&pb.OrderResponse{
			Id:         createdOrder.ID,
			Price:      float32(createdOrder.Price),
			Tax:        float32(createdOrder.Tax),
			FinalPrice: float32(createdOrder.FinalPrice),
		})
		if err != nil {
			return err
		}
	}
}

func (s *OrderService) ListOrders(context context.Context, request *pb.BlankRequest) (*pb.OrderListResponse, error) {
	order, err := s.GetAllOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var orders []*pb.OrderResponse
	for _, o := range order.Orders {
		orders = append(orders, &pb.OrderResponse{
			Id:    o.ID,
			Price: float32(o.Price),
			Tax:   float32(o.Tax),
		})
	}
	return &pb.OrderListResponse{
		Orders: orders,
	}, nil
}

func (s *OrderService) GetOrderById(context context.Context, request *pb.GetOrderByIdRequest) (*pb.OrderResponse, error) {
	order, err := s.GetOrderByIdUseCase.Execute(request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.OrderResponse{
		Id:         order.ID,
		Price:      float32(order.Price),
		Tax:        float32(order.Tax),
		FinalPrice: float32(order.FinalPrice),
	}, nil
}
