package main

import "net/http"

type handler struct {
	// handlers & gateway
}

func NewHander () *handler {
	return &handler{}
}

func(h *handler) handleCreateOrder(w http.ResponseWriter, r *http.Request){

}

func(h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/",h.handleCreateOrder)
}