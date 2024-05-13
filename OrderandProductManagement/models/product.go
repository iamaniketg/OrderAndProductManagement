package models

import "gorm.io/gorm"

// Product represents a product model
type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// SearchProducts searches for products based on the given query
func SearchProducts(query string) ([]Product, error) {
	var products []Product
	if err := db.Where("name LIKE ?", "%"+query+"%").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// CreateProduct creates a new product in the database
func CreateProduct(product *Product) error {
	return db.Create(product).Error
}

// GetAllProducts retrieves all products from the database
func GetAllProducts() ([]Product, error) {
	var products []Product
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// UpdateProduct updates a product in the database
func UpdateProduct(product Product) error {
	if err := db.Save(&product).Error; err != nil {
		return err
	}
	return nil
}

// DeleteProduct deletes a product from the database by its ID
func DeleteProduct(productID uint) error {
	// Retrieve the product by ID
	var product Product
	if err := db.First(&product, productID).Error; err != nil {
		return err // Product not found or other database error
	}

	// Delete the product
	if err := db.Delete(&product).Error; err != nil {
		return err // Error occurred while deleting the product
	}

	return nil // Product deleted successfully
}
