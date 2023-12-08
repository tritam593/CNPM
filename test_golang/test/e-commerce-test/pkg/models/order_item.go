package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/shopspring/decimal"
)

type OrderItem struct {
	ID              string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Order           Order
	OrderID         string `gorm:"size:36;index"`
	Product         Product
	ProductID       string `gorm:"size:36;index"`
	// Qty represents the quantity of the product in the order item.
	Qty             int
	BasePrice       decimal.Decimal `gorm:"type:decimal(16,2)"`
	BaseTotal       decimal.Decimal `gorm:"type:decimal(16,2)"`
	TaxAmount       decimal.Decimal `gorm:"type:decimal(16,2)"`
	TaxPercent      decimal.Decimal `gorm:"type:decimal(10,2)"`
	DiscountAmount  decimal.Decimal `gorm:"type:decimal(16,2)"`
	DiscountPercent decimal.Decimal `gorm:"type:decimal(10,2)"`
	SubTotal        decimal.Decimal `gorm:"type:decimal(16,2)"`
	// Sku represents the Stock Keeping Unit of the product in the order item.
	Sku             string          `gorm:"size:36;index"`
	Name            string          `gorm:"size:255"`
	Weight          decimal.Decimal `gorm:"type:decimal(10,2)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (o *OrderItem) BeforeCreate(db *gorm.DB) error {
	if o.ID == "" {
		o.ID = uuid.New().String()
	}

	return nil
}