package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Cart struct {
	ID             string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	CartItems      []CartItem
	UserID         string          `gorm:"size:36;index"`
	BaseTotalPrice decimal.Decimal `gorm:"type:decimal(16,2)"`
}

func (c *Cart) GetCart(db *gorm.DB, userID string) (*Cart, error) {
	var err error
	var cart Cart

	err = db.Debug().Preload("CartItems").
		Model(Cart{}).Where("User_ID = ?", userID).First(&cart).Error
	if err != nil {
		return nil, err
	}

	cart_item, _ := c.GetItems(db, cart.ID)
	cart.CartItems = cart_item

	return &cart, nil
}

func (c *Cart) CreateCart(db *gorm.DB) (*Cart, error) {

	err := db.Debug().Create(&c).Error
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Cart) CalculateCart(db *gorm.DB) (*Cart, error) {
	cartBaseTotalPrice := 0.0
	// Iterate through each cart item to calculate totals
	for _, item := range c.CartItems {
		itemBaseTotal, _ := item.BaseTotal.Float64()
		cartBaseTotalPrice += itemBaseTotal
	}

	var cart Cart

	c.BaseTotalPrice = decimal.NewFromFloat(cartBaseTotalPrice)
	err := db.Debug().First(&cart, "id = ?", c.ID).Updates(c).Error

	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (c *Cart) AddItem(db *gorm.DB, item CartItem) (*CartItem, error) {
	var existItem, updateItem CartItem
	var product Product
	// check the product already has in database
	err := db.Debug().Model(Product{}).Where("id = ?", item.ProductID).First(&product).Error
	if err != nil {
		return nil, err
	}

	basePrice, _ := product.Price.Float64()

	// Check if the item already exists in the cart
	err = db.Debug().Model(CartItem{}).
		Where("cart_id = ?", item.CartID).
		Where("product_id = ?", product.ID).
		First(&existItem).Error

	// If the item does not exist, create a new cart item and add it to the database
	if err != nil {
		item.BasePrice = product.Price
		item.BaseTotal = decimal.NewFromFloat(basePrice * float64(item.Qty))
		err = db.Debug().Create(&item).Error
		if err != nil {
			return nil, err
		}

		return &item, nil
	}

	// If the item already exists, update its quantity and recalculate totals
	updateItem.Qty = existItem.Qty + item.Qty
	updateItem.BaseTotal = decimal.NewFromFloat(basePrice * float64(updateItem.Qty))

	// Update the existing item in the database
	err = db.Debug().First(&existItem, "id = ?", existItem.ID).Updates(updateItem).Error
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (c *Cart) GetItems(db *gorm.DB, cartID string) ([]CartItem, error) {
	var items []CartItem

	err := db.Debug().Preload("Product").Model(&CartItem{}).
		Where("cart_id = ?", cartID).
		Order("created_at desc").
		Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (c *Cart) RemoveItemByID(db *gorm.DB, itemID string) (*CartItem, error) {
	var err error
	var item CartItem

	err = db.Debug().Model(&CartItem{}).Where("id = ?", itemID).First(&item).Error
	if err != nil {
		return nil, err
	}

	err = db.Debug().Unscoped().Delete(&item).Error
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// func (c *Cart) ClearCart(db *gorm.DB, cartID string) error {
// 	err := db.Debug().Where("cart_id = ?", cartID).Delete(&CartItem{}).Error
// 	if err != nil {
// 		return err
// 	}

// 	err = db.Debug().Where("id = ?", cartID).Delete(&Cart{}).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
