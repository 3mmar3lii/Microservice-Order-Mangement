package main

import "context"

func main() {

	// Store creation
	store := NewStore(context.Background())
	// Pass store to service layer
	svc :=NewService(store)
	// Excute Create order
	svc.CreateOrder(context.Background())
}