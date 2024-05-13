package handlers

import (
	"FirstProject/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ManageAdminRole handles requests to manage admin roles
func ManageAdminRole(c *fiber.Ctx) error {
	// Parse request body
	var requestData struct {
		UserID uint   `json:"user_id"`
		Role   string `json:"role"`
	}
	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	// Retrieve user by ID
	user, err := models.GetUserByID(requestData.UserID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
		})
	}

	// Update user's role
	user.Role = requestData.Role
	if err := models.UpdateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user role",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Admin role updated successfully",
		"user":    user,
	})
}

// Add product handler
func AddProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	// Create new product in the database
	err := models.CreateProduct(&product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to add product",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

// Remove product handler
func RemoveProduct(c *fiber.Ctx) error {
	// Parse product ID from request parameters or body
	productID := c.Params("id")

	// Convert the productID string to uint
	id, err := strconv.ParseUint(productID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid product ID",
			"error":   err.Error(),
		})
	}

	// Pass the converted product ID to the DeleteProduct function
	if err := models.DeleteProduct(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete product",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}

// Update product handler
func UpdateProduct(c *fiber.Ctx) error {
	// Parse product ID and updated product data from request body
	var updateData models.Product
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	// Update the product in the database
	err := models.UpdateProduct(updateData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update product",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Product updated successfully",
	})
}

// Order management handler
func OrderManagement(c *fiber.Ctx) error {
	// Implementation
	return c.SendString("Order management handler")
}

// Statistics generation handler
func GenerateStatistics(c *fiber.Ctx) error {
	// Implementation
	return c.SendString("Statistics generation handler")
}
