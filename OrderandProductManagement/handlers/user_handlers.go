package handlers

import (
	"FirstProject/models"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Signup handles user signup
func Signup(c *fiber.Ctx) error {
	// Parse request body to get user data
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	// Hash the user's password before saving to the database
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
			"error":   err.Error(),
		})
	}
	user.Password = hashedPassword

	// Create new user in the database
	err = models.CreateUser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// hashPassword hashes the provided password using bcrypt
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Login handles user login
func Login(c *fiber.Ctx) error {
	// Parse request body to get user login credentials
	var loginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	// Retrieve user record from the database using the provided email
	user, err := models.GetUserByEmail(loginReq.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	// Verify the provided password against the encrypted password stored in the database
	if !verifyPassword(user.Password, loginReq.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	// Generate a JWT token for the user
	token, err := GenerateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate authentication token",
			"error":   err.Error(),
		})
	}

	// Update the user's token field in the database
	user.Token = token
	if err := models.UpdateUserToken(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user token",
			"error":   err.Error(),
		})
	}

	// Return the authentication token in the response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func verifyPassword(encryptedPassword, plainTextPassword string) bool {
	// Compare the encrypted password with the plain-text password
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(plainTextPassword))
	return err == nil
}

var jwtSecret = []byte("your_secret_key")

// GenerateToken generates a JWT token for a user
func GenerateToken(userID uint) (string, error) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Generate encoded token
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Product search handler
func SearchProducts(c *fiber.Ctx) error {
	// Parse search query from request parameters or body
	searchQuery := c.Query("q")

	// Perform search query in the database to find matching products
	products, err := models.SearchProducts(searchQuery)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to search products",
			"error":   err.Error(),
		})
	}

	return c.JSON(products)
}

// Place order handler
func PlaceOrder(c *fiber.Ctx) error {
	// Parse order data from request body
	var orderData models.Order
	if err := c.BodyParser(&orderData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	// Place the order in the database
	err := models.PlaceOrder(&orderData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to place order",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Order placed successfully",
	})
}

// User dashboard handler
func UserDashboard(c *fiber.Ctx) error {
	// Retrieve user ID from request context or token
	userID, err := getUserIDFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve user ID",
			"error":   err.Error(),
		})
	}

	// Fetch user's orders from the database
	orders, err := models.GetOrdersByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch user orders",
			"error":   err.Error(),
		})
	}

	return c.JSON(orders)
}

// getUserIDFromContext retrieves the user ID from the request context or token
func getUserIDFromContext(c *fiber.Ctx) (uint, error) {
	// Retrieve JWT token from request headers or cookies
	token := c.Get("Authorization")

	// Validate and parse the JWT token to extract user ID
	userID, err := parseUserIDFromToken(token)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

// parseUserIDFromToken parses the user ID from the JWT token claims
func parseUserIDFromToken(token string) (uint, error) {
	// Implement JWT token parsing and claim extraction logic here
	// Example:
	claims, err := parseJWTToken(token)
	if err != nil {
		return 0, err
	}

	// Extract user ID claim from the token
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("user_id claim not found or invalid")
	}

	return uint(userID), nil
}

// parseJWTToken parses the JWT token and returns its claims
func parseJWTToken(tokenString string) (map[string]interface{}, error) {
	// Parse the JWT token string into a token object
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to extract claims")
	}

	return claims, nil
}
