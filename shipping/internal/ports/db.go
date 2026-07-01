package ports

import (
	"context"

	"github.com/CaioBernardelli/microservices/shipping/internal/application/core/domain"
)

type DBPort interface {
	Save(ctx context.Context, payment *domain.Shipping) error
	FindById(ctx context.Context, id int64) (domain.Shipping, error)
}





