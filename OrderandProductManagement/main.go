package main

import (
	"fmt"
	"log"

	"FirstProject/handlers"
	"FirstProject/middleware"
	"FirstProject/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Initialize database connection
	models.InitDB()

	// Middleware for logging
	app.Use(func(c *fiber.Ctx) error {
		fmt.Printf("Request: %s %s\n", c.Method(), c.Path())
		return c.Next()
	})

	// Middleware for authentication
	app.Use(middleware.AuthMiddleware)

	// User-side APIs
	app.Post("/signup", handlers.Signup)
	app.Post("/login", handlers.Login)
	app.Get("/products", handlers.SearchProducts)
	app.Post("/orders", handlers.PlaceOrder)
	app.Get("/dashboard", handlers.UserDashboard)

	// Admin-side APIs
	app.Put("/admin/roles", handlers.ManageAdminRole)
	app.Post("/admin/products", handlers.AddProduct)
	app.Delete("/admin/products/:id", handlers.RemoveProduct)
	app.Put("/admin/products/:id", handlers.UpdateProduct)
	app.Get("/admin/orders", handlers.OrderManagement)
	app.Get("/admin/statistics", handlers.GenerateStatistics)

	// Serve Swagger UI

	// Start the server
	port := ":3000"
	log.Printf("Server started on port %s", port)
	log.Fatal(app.Listen(port))

	// Initialize your Fiber app
	app = fiber.New()

	// Use the authentication middleware for all routes
	// app.Use(middleware.AuthMiddleware())

	// Define your API routes
	// For example:
	app.Get("/api/public", func(c *fiber.Ctx) error {
		return c.SendString("Public API - No authentication required")
	})

	// Start your Fiber app
	app.Listen(":3000")
}
