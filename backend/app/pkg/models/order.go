package models

import (
	"database/sql"
	"strconv"
	"strings"
	"app/pkg/consts"
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"

	"github.com/shopspring/decimal"
)

type Order struct {
	ID               string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	UserID           string `gorm:"size:36;index"`
	User             User
	OrderItems       []OrderItem
	OrderCustomer    *OrderCustomer
	Code             string `gorm:"size:50;index"`
	Status           int
	OrderDate        time.Time
	PaymentDue       time.Time
	PaymentStatus    string          `gorm:"size:50;index"`
	PaymentToken     sql.NullString  `gorm:"size:100;index"`
	BaseTotalPrice   decimal.Decimal `gorm:"type:decimal(16,2)"`
	ShippingCost     decimal.Decimal `gorm:"type:decimal(16,2)"`
	Note             string          `gorm:"type:text"`
	CancelledBy      sql.NullString  `gorm:"size:36"`
	CancelledAt      sql.NullTime
	CancellationNote sql.NullString `gorm:"size:255"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}

func (o *Order) BeforeCreate(db *gorm.DB) error {
	if o.ID == "" {
		o.ID = uuid.New().String()
	}

	o.Code = generateOrderNumber(db)

	return nil
}

func (o *Order) CreateOrder(db *gorm.DB) (*Order, error) {
	o.OrderDate = time.Now()
	o.PaymentDue = time.Now().AddDate(0, 0, 7)
	result := db.Debug().Create(&o)
	if result.Error != nil {
		return nil, result.Error
	}

	return o, nil
}

func (o *Order) FindByUserID(db *gorm.DB, UserID string) (*Order, error) {

	err := db.Debug().
		Preload("OrderCustomer").
		Preload("OrderItems").
		Preload("OrderItems.Product").
		Preload("User", func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Model(&Order{}).Where("User_ID = ?", UserID).
		First(&o).Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *Order) GetStatusLabel() string {
	var statusLabel string

	switch o.Status {
	case consts.OrderStatusPending:
		statusLabel = "PENDING"
	case consts.OrderStatusDelivered:
		statusLabel = "DELIVERED"
	case consts.OrderStatusReceived:
		statusLabel = "RECEIVED"
	case consts.OrderStatusCancelled:
		statusLabel = "CANCELLED"
	default:
		statusLabel = "UNKNOWN"
	}

	return statusLabel
}

func (o *Order) IsPaid() bool {
	return o.PaymentStatus == consts.OrderPaymentStatusPaid
}

func generateOrderNumber(db *gorm.DB) string {
	now := time.Now()
	month := now.Month()
	year := strconv.Itoa(now.Year())

	dateCode := "/ORDER/" + intToRoman(int(month)) + "/" + year

	var latestOrder Order

	err := db.Debug().Order("created_at DESC").Find(&latestOrder).Error

	latestNumber, _ := strconv.Atoi(strings.Split(latestOrder.Code, "/")[0])
	if err != nil {
		latestNumber = 1
	}

	number := latestNumber + 1

	invoiceNumber := strconv.Itoa(number) + dateCode

	return invoiceNumber
}

func intToRoman(num int) string {
	values := []int{
		1000, 900, 500, 400,
		100, 90, 50, 40,
		10, 9, 5, 4, 1,
	}

	symbols := []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV",
		"I"}
	roman := ""
	i := 0

	for num > 0 {
		// calculate the number of times this num is completly divisible by values[i]
		// times will only be > 0, when num >= values[i]
		k := num / values[i]
		for j := 0; j < k; j++ {
			// buildup roman numeral
			roman += symbols[i]

			// reduce the value of num.
			num -= values[i]
		}
		i++
	}
	return roman
}
