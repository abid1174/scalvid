# ScalVid Backend

A RESTful API backend service built with Go, featuring JWT authentication, PostgreSQL database integration, and a clean architecture design.

## 📋 Table of Contents

- [Overview](#overview)
- [Tech Stack](#tech-stack)
- [Architecture](#architecture)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Middleware](#middleware)
- [Configuration](#configuration)
- [Database](#database)
- [Authentication](#authentication)
- [Docker Setup](#docker-setup)

## 🎯 Overview

ScalVid Backend is a scalable REST API service designed to manage products with secure authentication. It follows clean architecture principles with clear separation of concerns, making it maintainable and testable.

## 🛠 Tech Stack

- **Language**: Go 1.23
- **Database**: PostgreSQL (latest)
- **Web Framework**: Standard library `net/http`
- **Database Driver**: `sqlx` + `pq` (PostgreSQL driver)
- **Configuration**: `godotenv` for environment variables
- **Containerization**: Docker & Docker Compose
- **Authentication**: Custom JWT implementation (HMAC-SHA256)

## 🏗 Architecture

The application follows a **layered architecture** pattern with dependency injection:

```
┌─────────────────────────────────────────┐
│          REST Layer (Handlers)          │
│  - HTTP Request/Response handling       │
│  - Route registration                   │
│  - Request validation                   │
└──────────────┬──────────────────────────┘
               │
┌──────────────▼──────────────────────────┐
│       Repository Layer (Data Access)    │
│  - Database operations                  │
│  - Data transformation                  │
└──────────────┬──────────────────────────┘
               │
┌──────────────▼──────────────────────────┐
│         Infrastructure Layer            │
│  - Database connections                 │
│  - External services (RabbitMQ, etc.)   │
└─────────────────────────────────────────┘
```

### Key Design Patterns

1. **Dependency Injection**: Dependencies are injected through constructors
2. **Repository Pattern**: Abstraction layer for data access
3. **Middleware Chain**: Composable middleware for cross-cutting concerns
4. **Interface Segregation**: Repository interfaces for loose coupling

## 📁 Project Structure

```
backend/
├── cmd/                          # Application commands
│   └── serve.go                  # Server initialization and wiring
│
├── config/                       # Configuration management
│   └── config.go                 # Environment variable loading
│
├── database/                     # Database models and mock data
│   ├── product.go                # Product model and mock operations
│   ├── user.go                   # User model
│   └── queries/                  # SQL migration scripts
│       └── 001-products.sql      # Product table schema
│
├── infra/                        # Infrastructure components
│   ├── db/
│   │   └── connection.go         # PostgreSQL connection setup
│   └── rabbitmq/                 # Message queue (future implementation)
│
├── repo/                         # Repository layer (data access)
│   ├── product.go                # Product repository implementation
│   └── user.go                   # User repository
│
├── rest/                         # REST API layer
│   ├── server.go                 # HTTP server setup
│   ├── handler/                  # Request handlers
│   │   ├── product/
│   │   │   ├── handler.go        # Product handler struct
│   │   │   ├── router.go         # Product route registration
│   │   │   ├── create_product.go # POST /products
│   │   │   ├── get_product.go    # GET /products
│   │   │   ├── get_product_by_id.go # GET /products/{id}
│   │   │   ├── update_product.go # PUT /products/{id}
│   │   │   └── delete_product.go # DELETE /products/{id}
│   │   └── user/                 # User handlers (future)
│   │
│   └── middleware/               # HTTP middleware
│       ├── middleware.go         # Middleware struct
│       ├── manage.go             # Middleware manager
│       ├── authentication_jwt.go # JWT authentication
│       ├── cors.go               # CORS handling
│       ├── logger.go             # Request logging
│       ├── preflight.go          # Preflight request handling
│       └── test.go               # Test middleware
│
├── utils/                        # Utility functions
│   ├── jwt.go                    # JWT token generation
│   └── response.go               # HTTP response helpers
│
├── docker-compose.yml            # Docker services configuration
├── Dockerfile                    # Multi-stage Docker build
├── go.mod                        # Go module dependencies
├── go.sum                        # Dependency checksums
└── main.go                       # Application entry point
```

## 🚀 Getting Started

### Prerequisites

- Go 1.23 or higher
- Docker & Docker Compose (for containerized deployment)
- PostgreSQL (if running locally without Docker)

### Environment Variables

Create a `.env` file in the backend directory:

```env
# Application
VERSION=1.0.0
HTTP_PORT=8000
GO_ENV=development

# Database
DB_HOST=localhost              # Use 'scalvid-db' in Docker
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=scalvid

# Authentication
JWT_SECRET_KEY=your_secret_key_here
```

### Local Development

1. **Install dependencies**:

   ```bash
   go mod download
   ```

2. **Set up PostgreSQL**:

   - Install PostgreSQL locally
   - Create database: `createdb scalvid`
   - Run migrations: `psql -d scalvid -f database/queries/001-products.sql`

3. **Run the application**:

   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:8000`

### Docker Deployment

1. **Start all services**:

   ```bash
   docker-compose up -d
   ```

   This will start:

   - `scalvid-backend`: Go application on port 8000
   - `scalvid-db`: PostgreSQL database on port 5432

2. **View logs**:

   ```bash
   docker-compose logs -f scalvid-backend
   ```

3. **Stop services**:
   ```bash
   docker-compose down
   ```

## 🔌 API Endpoints

All endpoints require JWT authentication via `Authorization: Bearer <token>` header.

### Products

| Method | Endpoint         | Description          | Auth Required |
| ------ | ---------------- | -------------------- | ------------- |
| POST   | `/products`      | Create a new product | ✅            |
| GET    | `/products`      | Get all products     | ✅            |
| GET    | `/products/{id}` | Get product by ID    | ✅            |
| PUT    | `/products/{id}` | Update a product     | ✅            |
| DELETE | `/products/{id}` | Delete a product     | ✅            |

### Request/Response Examples

**Create Product** (`POST /products`):

```json
{
  "title": "Product Name",
  "price": 99.99,
  "description": "Product description",
  "category": "Electronics",
  "imageUrl": "https://example.com/image.jpg"
}
```

**Response**:

```json
{
  "id": 1,
  "title": "Product Name",
  "price": 99.99,
  "description": "Product description",
  "category": "Electronics",
  "imageUrl": "https://example.com/image.jpg"
}
```

## 🔐 Middleware

The application uses a flexible middleware chain pattern:

### Global Middleware

Applied to all routes:

- **Preflight**: Handles OPTIONS requests for CORS
- **CORS**: Sets CORS headers (`Access-Control-Allow-*`)
- **Logger**: Logs HTTP requests with method, path, and duration

### Route-Level Middleware

Applied to specific routes:

- **JWT Authentication**: Validates JWT tokens using HMAC-SHA256

### Middleware Manager

The `middleware.Manager` provides two methods:

- `Use(middlewares...)`: Register global middleware
- `With(handler, middlewares...)`: Apply middleware to specific routes

Example:

```go
middlewareManager := middleware.NewManager()

// Global middleware
middlewareManager.Use(
    middleware.Preflight,
    middleware.Cors,
    middleware.Logger,
)

// Route-specific middleware
mux.Handle(
    "POST /products",
    middlewareManager.With(
        http.HandlerFunc(h.CreateProductHandler),
        h.middlewares.AuthenticationJWT,
    ),
)
```

## ⚙️ Configuration

Configuration is managed through environment variables using the `config` package.

### Config Structure

```go
type Config struct {
    Version      string  // Application version
    HttpPort     int     // HTTP server port
    GoEnv        string  // Environment (development/production)
    DbPort       string  // Database port
    DbUser       string  // Database username
    DbPassword   string  // Database password
    DbName       string  // Database name
    JwtSecretKey string  // JWT signing key
}
```

### Usage

```go
cfg := config.GetConfig()  // Singleton pattern
```

The configuration is loaded once at startup and cached for subsequent access.

## 💾 Database

### PostgreSQL Setup

The application uses PostgreSQL with the following configuration:

- **Driver**: `github.com/lib/pq`
- **ORM**: `github.com/jmoiron/sqlx` (enhanced database/sql)
- **Connection**: Singleton pattern via `db.NewConnection()`

### Schema

**Products Table**:

```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL CHECK (price >= 0),
    category_id VARCHAR(200) NOT NULL,
    image_url TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

### Repository Pattern

The repository pattern provides an abstraction layer over database operations:

```go
type ProductRepository interface {
    CreateProduct(product Product) (Product, error)
    GetProduct(id int) (Product, error)
    GetProducts() ([]Product, error)
    UpdateProduct(id int, product Product) (Product, error)
    DeleteProduct(id int) error
}
```

Benefits:

- Easy to mock for testing
- Decouples business logic from database implementation
- Allows swapping data sources without changing handlers

## 🔒 Authentication

The application uses **custom JWT (JSON Web Token)** implementation for authentication.

### JWT Structure

```
header.payload.signature
```

**Header**:

```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

**Payload**:

```json
{
  "sub": "user_id",
  "email": "user@example.com",
  "role": "admin",
  "exp": 1234567890,
  "iat": 1234567890
}
```

### Token Generation

```go
import "scalvid/utils"

payload := utils.Payload{
    Sub:   "user123",
    Email: "user@example.com",
    Role:  "admin",
}

token, err := utils.GenerateJWT(payload, secretKey)
```

### Token Validation

Tokens are validated in the `AuthenticationJWT` middleware:

1. Extract token from `Authorization: Bearer <token>` header
2. Split token into header, payload, and signature
3. Regenerate signature using secret key
4. Compare signatures (constant-time comparison recommended)
5. If valid, allow request; otherwise, return 401 Unauthorized

### Security Features

- **HMAC-SHA256**: Cryptographic signing algorithm
- **Token Expiration**: 1-hour expiry (configurable)
- **URL-safe Base64**: No padding encoding

## 🐳 Docker Setup

### Multi-Stage Build

The Dockerfile uses a multi-stage build for optimized image size:

1. **Builder Stage**:

   - Uses `golang:1.25.1-alpine`
   - Downloads dependencies
   - Compiles the Go binary

2. **Final Stage**:
   - Uses `alpine:latest` (minimal footprint)
   - Copies only the compiled binary
   - Exposes port 8000

### Docker Compose

The `docker-compose.yml` orchestrates two services:

**scalvid-backend**:

- Builds from local Dockerfile
- Depends on database (waits for health check)
- Auto-restarts unless stopped
- Maps port 8000

**scalvid-db**:

- PostgreSQL latest image
- Persistent volume for data
- Health check every 5 seconds
- Maps port 5432

### Commands

```bash
# Build and start
docker-compose up --build -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down

# Remove volumes (data)
docker-compose down -v

# Access database
docker exec -it scalvid-db psql -U your_db_user -d scalvid
```

## 📦 Dependencies

```go
require (
    github.com/joho/godotenv v1.5.1  // Environment variable loading
    github.com/jmoiron/sqlx v1.4.0   // Enhanced SQL library
    github.com/lib/pq v1.10.9        // PostgreSQL driver
)
```

## 🧪 Testing

(To be implemented)

Recommended testing strategy:

- **Unit tests**: Test handlers and repositories with mocks
- **Integration tests**: Test database operations with test database
- **E2E tests**: Test complete API flows

## 🚧 Future Enhancements

- [ ] User management endpoints
- [ ] Refresh token implementation
- [ ] Rate limiting middleware
- [ ] Request validation middleware
- [ ] RabbitMQ integration for async processing
- [ ] Comprehensive test suite
- [ ] API documentation (Swagger/OpenAPI)
- [ ] Pagination for GET /products
- [ ] Search and filtering
- [ ] Database migration tool (e.g., golang-migrate)

## 📄 License

(Add your license here)

## 👥 Contributors

(Add contributors here)

---

**Built with ❤️ using Go**
