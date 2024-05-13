package middleware

import (
	// "FirstProject/models"

	// "fmt"

	// "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	// "log"
	// "strconv"
)

// // AuthMiddleware is the authentication middleware
// func AuthMiddleware(c *fiber.Ctx) error {
// 	// Skip authentication for the signup and login routes
// 	if c.Path() == "/signup" || c.Path() == "/login" {
// 		return c.Next()
// 	}

// 	// Get token from request headers or cookies
// 	token := c.Get("Authorization")
// 	if token == "" {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"message": "Missing authentication token",
// 		})
// 	}

// 	// Validate token and extract user ID
// 	userID, err := models.ValidateToken(token)
// 	if err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"message": "Invalid authentication token",
// 		})
// 	}

// 	// Set user ID in locals for further processing
// 	c.Locals("userID", userID)

//		// Proceed to next middleware or handler
//		return c.Next()
//	}
//
// AuthMiddleware is the authentication middleware
func AuthMiddleware(c *fiber.Ctx) error {
	// Proceed to next middleware or handler
	return c.Next()
}
