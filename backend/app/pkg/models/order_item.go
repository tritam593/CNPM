package models

import (
	"time"

	"github.com/google/uuid"

	"github.com/shopspring/decimal"
)

type OrderItem struct {
	ID              string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	OrderID         string `gorm:"size:36;index"`
	Product         Product
	ProductID       string `gorm:"size:36;index"`
	Qty             int
	BasePrice       decimal.Decimal `gorm:"type:decimal(16,2)"`
	BaseTotal       decimal.Decimal `gorm:"type:decimal(16,2)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (o *OrderItem) BeforeCreate() error {
	if o.ID == "" {
		o.ID = uuid.New().String()
	}

	return nil
}