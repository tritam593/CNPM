package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderCustomer struct {
	ID        string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	User      User
	UserID    string `gorm:"size:36;index"`
	Order     Order
	OrderID   string `gorm:"size:36;index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o *OrderCustomer) BeforeCreate(db *gorm.DB) error {
	if o.ID == "" {
		o.ID = uuid.New().String()
	}

	return nil
}
