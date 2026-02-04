package models

import "time"

type OrderStatus string

const (
	OrderPending OrderStatus = "PENDING"
	OrderPaid    OrderStatus = "PAID"
	OrderFailed  OrderStatus = "FAILED"
)

type Order struct {
	ID        string      `json:"id"`
	UserID    string      `json:"user_id"`
	Status    OrderStatus `json:"status"`
	Amount    float64     `json:"amount"`
	Currency  string      `json:"currency"`
	Items     []OrderItem `json:"items"`
	CreatedAt time.Time   `json:"created_at"`
}

type OrderItem struct {
	ProductID string  `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
