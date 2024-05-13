# OrderAndProductManagement 🙏🙏
- Golang Rest API boilerplate built with GORM, Go-Fiber, and a MySQL database
-----
## Quickstart 🚀
#### To quickly get started with the Go-Fiber Boilerplate, follow these steps:
1. Clone the repository:
```
https://github.com/iamaniketg/OrderAndProductManagement.git
```

2. Install Dependencies:
```
cd go-fiber-boilerplate
go mod tidy
```
3. Connect the database: Create a ```.env``` file and put in a connection string:
```
dsn := "root:password@12345G@tcp(127.0.0.1:3306)/goDatabase?charset=utf8mb4&parseTime=True&loc=Local"
```
4. Start the project:
```
go run main.go
```
-----
## File Structure
```
OrderAndProductManagement/
├── main.go
├── handlers/
│   ├── user_handlers.go
│   └── admin_handlers.go
├── middleware/
│   └── auth_middleware.go
├── models/
│   └── database.go
│   └── user.go
│   └── product.go
│   └── order.go
│   └── customer.go
│   └── inventory.go
└── go.mod
```
-----
## Database
- The database folder initializes the database connection and registers the models. You can add new models by registering them in connect.go. The global DB variable is initialized in database.go.
-------
## Handlers
- The handlers folder defines each request for each model and how it interacts with the database. The functions are mapped by the router to the corresponding URL links.
-----
## Middleware
- The middleware folder contains a file for each middleware function. The security middleware is applied first to everything in router.go and applies general security middleware to the incoming requests. The JSON middleware serializes the incoming request so that it only allows JSON. Finally, the Authentication middleware is applied individually to requests that require the user to be logged in.
## Router
- The router file maps each incoming request to the corresponding function in handlers. It first applies the middleware and then groups the requests to each model and finally to the individual function.
-----
## Main.go 🚀
- The main.go file reads environment variables and applies the CORS middleware. You can change the allowed request sites in the configuration.
- Also contains api endpoints.
-----
## API Documentation
- Port := ":3000"
- This project provides a battery-included REST API that allows the user to interact with a MySQL database. The available endpoints are listed below:
1. User Endpoints:
```
- Post: /signup, /login, /orders
- Get: /products, /dashboard
```

2. Admin Endpoints:
```
- Put: /admin/roles, /admin/products/:id
- Post: /admin/products
- Get: /admin/orders, /admin/statistics
```
## License:
- This project is licensed under the MIT License.


