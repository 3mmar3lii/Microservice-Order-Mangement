package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ammar-alaa/common"
	pb "github.com/ammar-alaa/common/api"
	grpcutil "github.com/ammar-alaa/common/grpc"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHander(client pb.OrderServiceClient) *handler {
	return &handler{client}
}

func (h *handler) handleCreateOrder(w http.ResponseWriter, r *http.Request) {
	cusomerID := r.PathValue("customerID")
	var items []*pb.ItemsWithQuantity

	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := ValidateItems(items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	o, err := h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: cusomerID,
		Items:      items,
	})

	if grpcError := grpcutil.GRPCError(w, err); grpcError != nil {
		return
	}

	common.WriteJSON(w, http.StatusOK, o)
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.handleCreateOrder)
}

func ValidateItems(items []*pb.ItemsWithQuantity) error {

	if len(items) == 0 {
		return errors.New("items list cannot be empty")
	}

	for idx, item := range items {
		if item.ID == "" {
			return errors.New("item ID cannot be empty")
		}

		if item.Quantity <= 0 {
			return errors.New("item quantity must be greater than zero")
		}

		if item == nil {
			return fmt.Errorf("item at index %d cannot be nil", idx)
		}
	}

	return nil
}
