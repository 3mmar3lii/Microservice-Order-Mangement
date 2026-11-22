package main

import (
	"net/http"

	"github.com/ammar-alaa/common"
	pb "github.com/ammar-alaa/common/api"
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

	o,err := h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: cusomerID,
		Items:      items,
	})

	if err != nil {
		common.WriteError(w,http.StatusInternalServerError,err.Error())
	}

	common.WriteJSON(w,http.StatusOK,o)
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.handleCreateOrder)
}
