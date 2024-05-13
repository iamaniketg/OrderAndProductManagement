package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email" gorm:"unique"`
	Address string `json:"address"`
	Orders  []Order
}

// CreateCustomer creates a new customer in the database
func CreateCustomer(customer *Customer) error {
	return db.Create(customer).Error
}

// GetCustomerByEmail retrieves a customer from the database by email
func GetCustomerByEmail(email string) (*Customer, error) {
	var customer Customer
	if err := db.Where("email = ?", email).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}
