package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID            string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	CategoryID    string `gorm:"foreignKey:CategoryID"`
	ProductImages *ProductImage
	Category      *Category 
	// Sku represents the Stock Keeping Unit of the product and is indexed for efficient querying.
	Sku    string          `gorm:"size:100;index"`
	Name   string          `gorm:"size:255"`
	Price  decimal.Decimal `gorm:"type:decimal(16,2);"`
	Stock  int
	Weight decimal.Decimal `gorm:"type:decimal(10,2);"`
	// ShortDescription string          `gorm:"type:text"`
	Description string `gorm:"type:text"`
	Status      int    `gorm:"default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func (p *Product) GetProducts(db *gorm.DB) (*[]Product, error) {
	var err error
	var products []Product

	err = db.Debug().Model(&Product{}).Order("created_at desc").Find(&products).Error
	if err != nil {
		return nil, err
	}

	return &products, nil
}

func (p *Product) FindByID(db *gorm.DB, productID string) (*Product, error) {
	var err error
	var product Product

	err = db.Debug().Preload("Category").Preload("ProductImages").Model(&Product{}).Where("id = ?", productID).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *Product) CreateProduct(db *gorm.DB) error {
	var err error

	err = db.Debug().Create(&p).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *Product) DeleteProduct(db *gorm.DB, productID string) (*Product, error) {
	var pro Product
	db.Model(&Product{}).Where("ID=?", productID).Unscoped().Delete(&pro)
	return &pro, nil
}
