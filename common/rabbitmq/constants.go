package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// Queue Names (organized by service domain)
const (
	OrderCreatedQueue   = "order.created"
	OrderPaidQueue      = "order.paid"
)

// Exchange Names 
const (
	OrdersExchange   = "orders.exchange"
	PaymentsExchange = "payments.exchange"
)

// Routing Keys
const (
	RoutingKeyOrderCreated = "order.created"
	RoutingKeyOrderPaid    = "order.paid"
)

// Shared Queue Options (prevents PRECONDITION_FAILED mismatches)
type QueueConfig struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args 			 amqp.Table
}

var DefaultOrderCreatedConfig = QueueConfig{
	Name:       OrderCreatedQueue,
	Durable:    true,
	AutoDelete: false,
	Exclusive:  false,
	NoWait:     false,
}
