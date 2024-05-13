package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	ProductID uint `json:"product_id" gorm:"unique"`
	Quantity  int  `json:"quantity"`
}

// CreateInventory creates a new inventory record in the database
func CreateInventory(inventory *Inventory) error {
	return db.Create(inventory).Error
}

// UpdateInventory updates the quantity of an existing inventory record in the database
func UpdateInventory(productID uint, quantity int) error {
	return db.Model(&Inventory{}).Where("product_id = ?", productID).Update("quantity", quantity).Error
}

// GetInventoryByProductID retrieves the inventory record for a specific product from the database
func GetInventoryByProductID(productID uint) (*Inventory, error) {
	var inventory Inventory
	if err := db.Where("product_id = ?", productID).First(&inventory).Error; err != nil {
		return nil, err
	}
	return &inventory, nil
}
