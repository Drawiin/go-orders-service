package web_handler

import (
	"encoding/json"
	"net/http"

	"github.com/drawiin/go-orders-service/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type WebOrderHandler struct {
	CreateOrderUseCase  *usecase.CreateOrderUseCase
	GetAllOrders        *usecase.GetAllOrdersUseCase
	GetOrderByIdUseCase *usecase.GetOrderByIdUseCase
}

func NewWebOrderHandler(
	CreateOrderUseCase *usecase.CreateOrderUseCase,
	GetAllOrdersUseCase *usecase.GetAllOrdersUseCase,
	GetOrderByIdUseCase *usecase.GetOrderByIdUseCase,
) *WebOrderHandler {
	return &WebOrderHandler{
		CreateOrderUseCase:  CreateOrderUseCase,
		GetAllOrders:        GetAllOrdersUseCase,
		GetOrderByIdUseCase: GetOrderByIdUseCase,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.CreateOrderUseCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebOrderHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	output, err := h.GetAllOrders.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebOrderHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	output, err := h.GetOrderByIdUseCase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
