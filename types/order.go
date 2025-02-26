package types

import "time"

type OrderType string
type OrderSide string
type OrderStatus string

const (
	OrderTypeLimit  OrderType = "limit"
	OrderTypeMarket OrderType = "market"
	
	OrderSideBuy  OrderSide = "buy"
	OrderSideSell OrderSide = "sell"
	
	OrderStatusOpen      OrderStatus = "open"
	OrderStatusClosed    OrderStatus = "closed"
	OrderStatusCanceled  OrderStatus = "canceled"
	OrderStatusRejected  OrderStatus = "rejected"
)

type Order struct {
	ID            string
	ClientOrderID string
	Symbol        string
	Type          OrderType
	Side          OrderSide
	Status        OrderStatus
	Price         float64
	Amount        float64
	Filled        float64
	Remaining     float64
	Cost          float64
	Created       time.Time
	Updated       time.Time
} 