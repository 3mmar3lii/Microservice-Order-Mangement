package main

import (
	"errors"
	"net/http"

	"github.com/ammar-alaa/common"
	pb "github.com/ammar-alaa/common/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		common.WriteError(w,http.StatusBadRequest,err.Error())
		return 
	}

	o, err := h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: cusomerID,
		Items:      items,
	})

	// Handling gRPC error 
if err != nil {
    rStatus := status.Convert(err)
    if rStatus.Code() == codes.InvalidArgument {
        common.WriteError(w, http.StatusBadRequest, rStatus.Message())
        return
    }
    common.WriteError(w, http.StatusInternalServerError, err.Error())
    return
}

	common.WriteJSON(w, http.StatusOK, o)
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.handleCreateOrder)
}

func ValidateItems(items []*pb.ItemsWithQuantity) error {

	if len(items) == 0 {
		return errors.New("")
	}

	for _, i := range items {
		if i.ID == "" {
			return errors.New("")
		}
		if i.Quantity <= 0 {
			return errors.New("")
		}
	}

	return nil
}
