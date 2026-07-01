package ports

import (
	"context"

	"github.com/CaioBernardelli/microservices/shipping/internal/application/core/domain"
)

type APIPort interface {
	Create(ctx context.Context, shipping domain.Shipping) (domain.Shipping, error)
}