# Scalvid Backend Architecture Plan

## Current Architecture: Layered / Repository Pattern

```
main.go
  └── cmd/          (Composition Root / Bootstrap)
        └── rest/         (Transport / HTTP Layer)
              └── handler/     (Request Handling)
                    └── repo/       (Data Access / Repository)
                          └── infra/      (Infrastructure / DB Connection)
```

### Layers

| Layer | Package | Responsibility |
|---|---|---|
| **Entrypoint** | `main.go` | Starts the app |
| **Bootstrap** | `cmd/` | Wires all dependencies together (composition root) |
| **Config** | `config/` | Loads environment variables |
| **Transport** | `rest/` | HTTP server setup, global middleware |
| **Handlers** | `rest/handler/` | HTTP handlers (one sub-package per resource) |
| **Middleware** | `rest/middleware/` | Auth (JWT), CORS, logging |
| **Repository** | `repo/` | Data access with interface + concrete implementation |
| **Infrastructure** | `infra/db/` | Database connection management |
| **Utils** | `utils/` | Shared helpers (JWT, HTTP responses) |

### Key Patterns

- **Repository Pattern** — Interface + private struct implementation, decouples handlers from the database.
- **Manual Dependency Injection** — `cmd/serve.go` acts as the composition root, constructing and wiring all dependencies top-down.
- **Resource-based handler grouping** — Each resource has its own sub-package with separate files per operation and a `router.go` for route registration.
- **Middleware chain** — Custom middleware manager supports global and per-route middleware.

---

## Phase 1: Add Service/Use-Case Layer

### Why

Currently handlers talk directly to repositories. Adding a service layer provides a place for business logic and enables cross-module communication.

### Inter-Module Communication: Narrow Interface Approach

When modules need each other (e.g., product needs order data AND order needs product data), use the **narrow interface pattern** to avoid import cycles.

**Rule: the consumer defines the interface for what it needs.**

#### Product Service

```go
// service/product/service.go
package product

import "scalvid/repo"

// What product needs from orders — defined HERE, not in the order package
type OrderChecker interface {
    HasActiveOrders(productID int) (bool, error)
}

type Service interface {
    GetProduct(id int) (*repo.Product, error)
    GetProducts() ([]*repo.Product, error)
    GetProductPrice(id int) (float64, error)
    DeleteProduct(id int) error
}

type service struct {
    productRepo  repo.ProductRepository
    orderChecker OrderChecker
}

func NewService(pr repo.ProductRepository) Service {
    return &service{productRepo: pr}
}

func (s *service) SetOrderChecker(oc OrderChecker) {
    s.orderChecker = oc
}

func (s *service) GetProductPrice(id int) (float64, error) {
    p, err := s.productRepo.GetProduct(id)
    if err != nil {
        return 0, err
    }
    return p.Price, nil
}

func (s *service) DeleteProduct(id int) error {
    hasOrders, err := s.orderChecker.HasActiveOrders(id)
    if err != nil {
        return err
    }
    if hasOrders {
        return errors.New("cannot delete product with active orders")
    }
    return s.productRepo.DeleteProduct(id)
}
```

#### Order Service

```go
// service/order/service.go
package order

import "scalvid/repo"

// What order needs from products — defined HERE, not in the product package
type ProductPricer interface {
    GetProductPrice(id int) (float64, error)
}

type Service interface {
    CreateOrder(order CreateOrderInput) (*repo.Order, error)
    HasActiveOrders(productID int) (bool, error)
}

type service struct {
    orderRepo     repo.OrderRepository
    productPricer ProductPricer
}

func NewService(or repo.OrderRepository, pp ProductPricer) Service {
    return &service{orderRepo: or, productPricer: pp}
}

func (s *service) CreateOrder(input CreateOrderInput) (*repo.Order, error) {
    price, err := s.productPricer.GetProductPrice(input.ProductID)
    if err != nil {
        return nil, err
    }
    order := repo.Order{
        ProductID:  input.ProductID,
        Quantity:   input.Quantity,
        TotalPrice: price * float64(input.Quantity),
    }
    return s.orderRepo.CreateOrder(order)
}

func (s *service) HasActiveOrders(productID int) (bool, error) {
    orders, err := s.orderRepo.GetOrdersByProductID(productID)
    if err != nil {
        return false, err
    }
    return len(orders) > 0, nil
}
```

#### Wiring in Composition Root

```go
// cmd/serve.go
func Serve() {
    config := config.GetConfig()
    dbConn := db.NewConnection(config.DB)

    productRepo := repo.NewProductRepository(dbConn)
    orderRepo   := repo.NewOrderRepository(dbConn)

    // Step 1: create product service (without order dependency)
    productSvc := product.NewService(productRepo)

    // Step 2: create order service (inject product service as ProductPricer)
    orderSvc := order.NewService(orderRepo, productSvc)

    // Step 3: wire order service back into product as OrderChecker
    productSvc.SetOrderChecker(orderSvc)

    productHandler := productHandler.NewHandler(middlewares, productSvc)
    orderHandler   := orderHandler.NewHandler(middlewares, orderSvc)

    server := rest.NewServer(config, productHandler, orderHandler)
    server.StartServer()
}
```

#### How It Works

```
product package                          order package
┌─────────────────────┐                 ┌─────────────────────┐
│                     │                 │                     │
│  OrderChecker {     │                 │  ProductPricer {    │
│    HasActiveOrders  │◄── satisfied by ──  Service           │
│  }                  │                 │  }                  │
│                     │                 │                     │
│  Service            ├── satisfies ──► │  GetProductPrice    │
│    GetProductPrice  │                 │                     │
│    DeleteProduct    │                 │  Service            │
│                     │                 │    HasActiveOrders  │
│                     │                 │    CreateOrder      │
└─────────────────────┘                 └─────────────────────┘
```

- Neither package imports the other.
- Go's implicit interface satisfaction (duck typing) makes this work.
- Easy to test — just mock the 1-method interface.

---

## Phase 2: Modular Monolith

Reorganize so each module is fully self-contained:

```
backend/
├── cmd/
│   └── serve.go
├── config/
├── infra/
│
├── modules/
│   ├── product/
│   │   ├── handler/
│   │   ├── service/
│   │   ├── repo/
│   │   └── model/
│   │
│   └── order/
│       ├── handler/
│       ├── service/
│       ├── repo/
│       └── model/
│
├── contract/                     # shared interfaces
│   ├── product.go
│   └── order.go
│
├── pkg/                          # shared utilities only
│   ├── middleware/
│   ├── response/
│   └── jwt/
│
└── main.go
```

### Key Rules

1. **Modules never import each other** — they communicate through `contract/` interfaces.
2. **Each module owns its data** — no cross-module SQL joins.
3. **Shared contract package** — interfaces that both modules depend on.

### Shared Contracts

```go
// contract/product.go
package contract

import "context"

type ProductPricer interface {
    GetProductPrice(ctx context.Context, id int) (float64, error)
}

type ProductInfo struct {
    ID    int
    Title string
    Price float64
}
```

```go
// contract/order.go
package contract

import "context"

type OrderChecker interface {
    HasActiveOrders(ctx context.Context, productID int) (bool, error)
}
```

### Data Ownership

```go
// BAD — order repo reaches into products table
query := `
    INSERT INTO orders (product_id, quantity, total_price)
    SELECT $1, $2, p.price * $2
    FROM products p WHERE p.id = $1
`

// GOOD — order service asks product service for the price
price, err := s.productPricer.GetProductPrice(ctx, input.ProductID)
order := Order{
    ProductID:  input.ProductID,
    Quantity:   input.Quantity,
    TotalPrice: price * float64(input.Quantity),
}
```

---

## Phase 3: Microservices

Each module becomes its own deployable service. The service layer code does not change — only the wiring and adapter implementations change.

### Adapter Pattern: Swap Function Call for Network Call

```go
// adapter/grpc_product_pricer.go
package adapter

import (
    "scalvid/contract"
    pb "scalvid/proto/product"
)

type GRPCProductPricer struct {
    client pb.ProductServiceClient
}

func (g *GRPCProductPricer) GetProductPrice(ctx context.Context, id int) (float64, error) {
    resp, err := g.client.GetPrice(ctx, &pb.GetPriceRequest{ProductId: int64(id)})
    if err != nil {
        return 0, err
    }
    return resp.Price, nil
}
```

### Microservice Wiring

```go
// In order microservice's main.go
func main() {
    productConn := grpc.Dial("product-service:50051")
    productPricer := adapter.NewGRPCProductPricer(productConn)

    orderSvc := order.NewService(orderRepo, productPricer) // same constructor
}
```

### Final Directory Structure

```
scalvid/
├── contract/                        # shared proto/interfaces (separate go module)
│   ├── product.go
│   └── order.go
│
├── product-service/                 # standalone binary
│   ├── cmd/main.go
│   ├── handler/
│   ├── service/
│   ├── repo/
│   ├── adapter/
│   │   └── grpc_order_checker.go
│   └── Dockerfile
│
├── order-service/                   # standalone binary
│   ├── cmd/main.go
│   ├── handler/
│   ├── service/
│   ├── repo/
│   ├── adapter/
│   │   └── grpc_product_pricer.go
│   └── Dockerfile
│
└── docker-compose.yml
```

---

## Migration Summary

| Concern | Monolith | Modular Monolith | Microservices |
|---|---|---|---|
| **Communication** | Direct function call | Direct function call | gRPC / HTTP / Queue |
| **Interface location** | Each module's package | Shared `contract/` package | Shared proto/contract repo |
| **Implementation** | Real service struct | Real service struct | Adapter (gRPC/HTTP client) |
| **Database** | Shared DB | Separate schemas, same DB | Separate DBs |
| **Deployment** | Single binary | Single binary | Separate binaries |
| **Service code changes** | — | Nothing | **Nothing** |

**Core principle: business logic (service layer) never changes. Only the wiring in the composition root and adapter implementations change.**
