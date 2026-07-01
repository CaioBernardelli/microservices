package db

import (
	"context"
	"fmt"

	"github.com/CaioBernardelli/microservices/shipping/internal/application/core/domain"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Shipping struct {
	gorm.Model
	OrderID      int64
	DeliveryDays int32
}

type Adapter struct {
	db *gorm.DB
}

func (a Adapter) FindById(ctx context.Context, id int64) (domain.Shipping, error) {
	var shippingEntity Shipping

	res := a.db.WithContext(ctx).First(&shippingEntity, id)

	shipping := domain.Shipping{
		ID:           int64(shippingEntity.ID),
		OrderId:      shippingEntity.OrderID,
		DeliveryDays: shippingEntity.DeliveryDays,
		CreatedAt:    shippingEntity.CreatedAt.Unix(),
	}

	return shipping, res.Error
}

func (a Adapter) Save(ctx context.Context, shipping *domain.Shipping) error {

	shippingModel := Shipping{
		OrderID:      shipping.OrderId,
		DeliveryDays: shipping.DeliveryDays,
	}

	res := a.db.WithContext(ctx).Create(&shippingModel)

	if res.Error == nil {
		shipping.ID = int64(shippingModel.ID)
	}

	return res.Error
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {

	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})

	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	if err := db.Use(otelgorm.NewPlugin(otelgorm.WithDBName("shipping"))); err != nil {
		return nil, fmt.Errorf("db otel plugin error: %v", err)
	}

	err := db.AutoMigrate(&Shipping{})

	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}