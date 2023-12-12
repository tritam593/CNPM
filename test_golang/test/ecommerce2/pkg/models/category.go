package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	// ParentID  string `gorm:"size:36;"`
	Products  []Product `gorm:"foreignKey:CategoryID"`
	Name      string    `gorm:"size:100;"`
	Slug      string    `gorm:"size:100;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Category) CreateCategory(db *gorm.DB) error {
	var err error

	err = db.Debug().Create(&c).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *Category) DeleteCategory(db *gorm.DB, CategoryID string) (*Category, error) {
	var cat Category
	db.Model(&Category{}).Where("ID=?", CategoryID).Unscoped().Delete(&cat)
	return &cat, nil
}

func (c *Category) GetCategoryByID(db *gorm.DB, CategoryID string) (*Category, error) {
	var cat Category
	db.Debug().Preload("Products").Find(&cat).Where("ID=?", CategoryID)
	return &cat, nil
}

func (c *Category) GetCategories(db *gorm.DB) (*[]Category, error) {
	var cat []Category
	db.Debug().Find(&cat)
	return &cat, nil
}


func (c *Category) UpdateCategory(db *gorm.DB, catID string) (*Category, error) {
	db.Debug().Model(&Category{}).Where("id = ?", catID).Updates(&c)
	return c.GetCategoryByID(db, catID)
}


