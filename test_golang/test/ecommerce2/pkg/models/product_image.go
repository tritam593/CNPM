package models

import (
	"time"
	"gorm.io/gorm"
)

type ProductImage struct {
	ID         string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Product    Product
	ProductID  string `gorm:"size:36;index"`
	Path       string `gorm:"type:text"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (img *ProductImage) CreateProductImg(db *gorm.DB) error {
	var err error

	err = db.Debug().Create(&img).Error
	if err != nil {
		return err
	}

	return nil
}

func (img *ProductImage) DeleteProductImage(db *gorm.DB, ProductImageID string) (*ProductImage,error) {
	var p_img ProductImage
	db.Model(&ProductImage{}).Where("ID=?", ProductImageID).Unscoped().Delete(&p_img)
	return &p_img, nil
}
