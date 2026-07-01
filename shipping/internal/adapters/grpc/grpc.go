package grpc

import (
	"context"

	"github.com/CaioBernardelli/microservices/shipping/internal/application/core/domain"
	"github.com/CaioBernardelli/microservices-proto/golang/shipping"
	log "github.com/sirupsen/logrus"
)

func (a Adapter) Create(ctx context.Context, request *shipping.CreateShippingRequest) (*shipping.CreateShippingResponse, error) {
	log.WithContext(ctx).Info("Creating shipping...")

	// Soma a quantidade total de unidades
	totalItems := 0
	for _, item := range request.Items {
		totalItems += int(item.Quantity)
	}

	// Prazo mínimo de 1 dia
	deliveryDays := int32(1)

	// A cada 5 unidades adiciona 1 dia
	if totalItems > 5 {
		deliveryDays += int32(totalItems / 5)
	}

	newShipping := domain.NewShipping(
		request.OrderId,
		deliveryDays,
	)

	result, err := a.api.Create(ctx, newShipping)
	if err != nil {
		return nil, err
	}

	return &shipping.CreateShippingResponse{
		DeliveryDays: result.DeliveryDays,
	}, nil
}