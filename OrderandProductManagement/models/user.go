package models

import (
	"fmt"
	"log"
	"strconv"

	// "time"
	// "crypto/rand"
	// "encoding/base64"
	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtSecret = []byte("your_secret_key")

// // GenerateToken generates a JWT token for a user/customer
// func GenerateToken(userID uint) (string, error) {
// 	// Create a new JWT token
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	// Set claims
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["userID"] = userID
// 	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

// 	// Generate encoded token
// 	tokenString, err := token.SignedString(jwtSecret)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }

// ValidateToken validates a JWT token and returns the user/customer ID if valid
func ValidateToken(tokenString string) (uint, error) {
	// Log the start of the token validation process
	log.Println("Starting token validation...")

	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Log the start of token parsing
		log.Println("Parsing token...")

		// Check if the signing method is valid
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// Log unexpected signing method
			log.Printf("Unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Log successful token parsing
		log.Println("Token parsed successfully.")

		// Return the JWT secret for validation
		return jwtSecret, nil
	})
	if err != nil {
		// Log token parsing error
		log.Printf("Error parsing token: %v", err)
		return 0, err
	}

	// Log token validation
	log.Println("Validating token...")

	// Check if the token is valid
	if !token.Valid {
		// Log invalid token
		log.Println("Invalid token.")
		return 0, jwt.ErrInvalidKey
	}

	// Log successful token validation
	log.Println("Token validated successfully.")

	// Extract user/customer ID from claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		// Log invalid token claims
		log.Println("Invalid token claims.")
		return 0, jwt.ErrInvalidKey
	}

	// Parse the user ID from the claims
	userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["userID"]), 10, 64)
	if err != nil {
		// Log error parsing user ID
		log.Printf("Error parsing user ID: %v", err)
		return 0, err
	}

	// Log successful user ID extraction
	log.Printf("User ID extracted successfully: %d", userID)

	// Return the validated user ID
	return uint(userID), nil
}

// GetUserByID retrieves a user by ID from the database
func GetUserByID(userID uint) (*User, error) {
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err // User not found
		}
		return nil, err // Other database error
	}
	return &user, nil
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Token    string
}

// UpdateUser updates a user in the database
func UpdateUser(user *User) error {
	return db.Save(user).Error
}

// CreateUser creates a new user in the database
func CreateUser(user *User) error {
	return db.Create(user).Error
}

// GetUserByEmail retrieves a user from the database by email
func GetUserByEmail(email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func AuthenticateUser(email, password string) (*User, error) {
	// Retrieve user from database by email
	user, err := GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	// Compare the provided password with the hashed password stored in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err // Passwords do not match
	}

	// Passwords match, return the user
	return user, nil
}

// UpdateUserToken updates the token field of the user record in the database
func UpdateUserToken(user *User) error {
	// Update the user record in the database to store the generated token
	if err := db.Model(&User{}).Where("id = ?", user.ID).Update("token", user.Token).Error; err != nil {
		return err
	}
	return nil
}
