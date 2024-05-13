package models

import "gorm.io/gorm"

// Order represents an order model
type Order struct {
	gorm.Model
	UserID     uint    `json:"user_id"`
	ProductID  uint    `json:"product_id"`
	Quantity   uint    `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

// PlaceOrder places an order in the database
func PlaceOrder(order *Order) error {
	// Create the order record in the database
	if err := db.Create(order).Error; err != nil {
		return err
	}
	return nil
}

// CreateOrder creates a new order in the database
func CreateOrder(order *Order) error {
	return db.Create(order).Error
}

// GetOrdersByUserID retrieves orders for a specific user from the database
func GetOrdersByUserID(userID uint) ([]Order, error) {
	var orders []Order
	if err := db.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
