package models

import "time"

type CustomerInfo struct {
	CustomerID string `json:"customer_id"`
	FullName   string `json:"full_name"`
	Email      string `json:"email"`
	Tier       string `json:"tier"`
}

type OrderLineItem struct {
	ProductID   string  `json:"product_id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	TotalAmount float64 `json:"total_amount"`
}

type OrderEvent struct {
	OrderID      string          `json:"order_id"`
	TrackingID   string          `json:"tracking_id"`
	Customer     CustomerInfo    `json:"customer"`
	Items        []OrderLineItem `json:"items"`
	Subtotal     float64         `json:"subtotal"`
	Tax          float64         `json:"tax"`
	Total        float64         `json:"total"`
	Currency     string          `json:"currency"`
	ShippingCity string          `json:"shipping_city"`
	CreatedAt    time.Time       `json:"created_at"`
}
