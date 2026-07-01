package domain

import "time"

type Shipping struct {
	ID           int64 `json:"id"`
	OrderId      int64 `json:"order_id"`
	DeliveryDays int32 `json:"delivery_days"`
	CreatedAt    int64 `json:"created_at"`
}

func NewShipping(orderId int64, deliveryDays int32) Shipping {
	return Shipping{
		OrderId:      orderId,
		DeliveryDays: deliveryDays,
		CreatedAt:    time.Now().Unix(),
	}
}