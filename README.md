# Go Fiber JWT 
#### Overview
Learning Fiber and JWT to create efficient and secure web applications in Go. With Fiber's speed and simplicity, combined with JWT's authentication capabilities, you can build modern and protected web services, making you a valuable developer in today's digital landscape.

### Features
- Fiber, JWT middlewares, security bcrypt
- RESTful routes
- User Product models
- Postgres Database connection, migration

### Prerequisites
Before you begin, ensure you have met the following requirements:

Go 1.16+ installed
Fiber v2.49.2
JWT middleware for Fiber v1.0.7
JWT library v5.0.0
godotenv v1.5.1

# API Documentation

## Middleware
  - Get("/", handler.Hello)

## Auth
 - Post("/login", handler.Login)

## Product
  - Get("/", handler.GetAllProducts)
  - Get("/:id", handler.GetProduct)
  - Post("/", middleware.Protected(), handler.CreateProduct)
  - Delete("/:id", middleware.Protected(), handler.DeleteProduct)
## USER
  - Get("/:id", handler.GetUser)
  - Post("/", handler.CreateUser)
  - Patch("/:id", middleware.Protected(), handler.UpdateUser)
  - Delete("/:id", middleware.Protected(), handler.DeleteUser)

## Product
  - Get("/", handler.GetAllProducts)
  - Get("/:id", handler.GetProduct)
  - Post("/", middleware.Protected(), handler.CreateProduct)
  - Delete("/:id", middleware.Protected(), handler.DeleteProduct)
