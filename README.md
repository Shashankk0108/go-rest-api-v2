<div align="center">

# Go REST API v2.0

### Production-Ready RESTful API with Clean Architecture

[![Go Version](https://img.shields.io/badge/Go-1.23.5-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![Gin Framework](https://img.shields.io/badge/Gin-v1.10.0-00ADD8?style=for-the-badge)](https://gin-gonic.com/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-316192?style=for-the-badge&logo=postgresql)](https://www.postgresql.org/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker)](https://www.docker.com/)
[![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)](LICENSE)

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=for-the-badge)](CONTRIBUTING.md)
[![Code Quality](https://img.shields.io/badge/code%20quality-A+-brightgreen?style=for-the-badge)]()
[![Maintained](https://img.shields.io/badge/Maintained-Yes-green.svg?style=for-the-badge)]()
[![Clean Architecture](https://img.shields.io/badge/architecture-clean-blue?style=for-the-badge)]()
[![REST API](https://img.shields.io/badge/API-RESTful-orange?style=for-the-badge)]()
[![JWT Auth](https://img.shields.io/badge/auth-JWT-red?style=for-the-badge)]()

[Features](#-features) • [Architecture](#-architecture) • [Quick Start](#-quick-start) • [API Documentation](#-api-documentation) • [Deployment](#-deployment)

</div>

---

## Overview

A **production-ready** RESTful API built with Go, following clean architecture principles and industry best practices. This project showcases professional-grade API development with comprehensive authentication, authorization, CRUD operations, and database management.

### What Makes This Special?

- **Clean Architecture**: Organized in layers (handler → service → repository) for maintainability
- **JWT Authentication**: Secure token-based authentication with middleware
- **PostgreSQL Integration**: Robust database operations with GORM
- **Docker Ready**: Complete containerization with Docker Compose
- **API Documentation**: Swagger/OpenAPI documentation
- **Production Standards**: Error handling, logging, CORS, validation
- **Task Automation**: Makefile and Taskfile for common operations

---

## Features

### Core Functionality

- **User Management**
  - User registration with password hashing (bcrypt)
  - JWT-based authentication
  - User profile management (CRUD operations)
  - Soft delete support

- **Product Management**
  - Create, read, update, delete products
  - Product categorization
  - User-specific product listings
  - Pagination support

### Technical Features

- **Security**
  - JWT token authentication
  - Password hashing with bcrypt
  - CORS middleware
  - Input validation
  - SQL injection prevention (GORM)

- **Database**
  - PostgreSQL with GORM ORM
  - Auto-migrations
  - Relationship management
  - Transaction support

- **Developer Experience**
  - Swagger API documentation
  - Postman collection included
  - Docker containerization
  - Hot reload in development
  - Comprehensive logging

---

## Architecture

### Project Structure

```
go-rest-api-v2/
├── main.go                      # Application entry point
├── internal/
│   ├── config/                  # Configuration management
│   │   └── config.go
│   ├── database/                # Database connection & migrations
│   │   └── database.go
│   ├── models/                  # Domain models
│   │   ├── user.go
│   │   └── product.go
│   ├── repository/              # Data access layer
│   │   ├── user_repository.go
│   │   └── product_repository.go
│   ├── service/                 # Business logic layer
│   │   ├── auth_service.go
│   │   ├── user_service.go
│   │   └── product_service.go
│   ├── handler/                 # HTTP handlers (controllers)
│   │   ├── auth_handler.go
│   │   ├── user_handler.go
│   │   └── product_handler.go
│   ├── middleware/              # HTTP middleware
│   │   ├── auth.go
│   │   ├── cors.go
│   │   └── logger.go
│   └── utils/                   # Utility functions
│       ├── jwt.go
│       └── response.go
├── api/                         # API documentation & examples
│   ├── postman_collection.json
│   └── examples.md
├── docker-compose.yml           # Docker services configuration
├── Dockerfile                   # Application container
├── Taskfile.yml                 # Task automation
├── Makefile                     # Alternative task automation
├── .env.example                 # Environment variables template
├── LICENSE                      # MIT License
├── CONTRIBUTING.md              # Contribution guidelines
└── README.md                    # This file

```

### Architecture Diagram

```
┌─────────────────────────────────────────────────────────┐
│                     HTTP Clients                        │
│              (Browser, Postman, Mobile App)             │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│                  Middleware Layer                       │
│          (CORS, Auth, Logging, Recovery)                │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│                  Handler Layer                          │
│        (AuthHandler, UserHandler, ProductHandler)       │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│                  Service Layer                          │
│         (Business Logic & Validation)                   │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│                Repository Layer                         │
│            (Data Access & ORM)                          │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│                PostgreSQL Database                      │
└─────────────────────────────────────────────────────────┘
```

### Design Patterns

- **Dependency Injection**: Services and repositories are injected
- **Repository Pattern**: Abstract data access logic
- **Service Layer Pattern**: Encapsulate business logic
- **Middleware Pattern**: Cross-cutting concerns (auth, logging)
- **Factory Pattern**: Create instances of repositories and services

---

## Prerequisites

Before you begin, ensure you have the following installed:

| Tool | Version | Purpose |
|------|---------|---------|
| **Go** | 1.23.5+ | Programming language |
| **Docker** | 20.10+ | Containerization |
| **Docker Compose** | 2.0+ | Multi-container orchestration |
| **PostgreSQL** | 16+ | Database (or use Docker) |
| **Task** | 3.0+ | Task automation (optional) |
| **Make** | 4.0+ | Build automation (optional) |

### Installation Links

- [Install Go](https://golang.org/doc/install)
- [Install Docker](https://docs.docker.com/get-docker/)
- [Install Docker Compose](https://docs.docker.com/compose/install/)
- [Install Task](https://taskfile.dev/installation/)

---

## Quick Start

### Option 1: Using Docker (Recommended)

1. **Clone the repository**
   ```bash
   git clone https://github.com/botsgalaxy/go-rest-api-v2.git
   cd go-rest-api-v2
   ```

2. **Start the application**
   ```bash
   docker-compose up -d
   ```

3. **Verify it's running**
   ```bash
   curl http://localhost:8080/health
   ```

   You should see:
   ```json
   {
     "status": "ok",
     "message": "Server is running"
   }
   ```

### Option 2: Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/botsgalaxy/go-rest-api-v2.git
   cd go-rest-api-v2
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

4. **Start PostgreSQL** (if not using Docker)
   ```bash
   # Using Docker for just the database
   docker run --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=go_rest_api -p 5432:5432 -d postgres:16-alpine
   ```

5. **Run the application**
   ```bash
   go run main.go
   ```

   Or using Task/Make:
   ```bash
   task dev
   # or
   make dev
   ```

---

## Configuration

### Environment Variables

Create a `.env` file in the root directory (use `.env.example` as template):

```env
# Server Configuration
SERVER_HOST=0.0.0.0
SERVER_PORT=8080
GIN_MODE=debug  # Use 'release' in production

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=go_rest_api
DB_SSLMODE=disable

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
JWT_EXPIRATION=24h

# CORS
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:8080
```

---

## API Documentation

### Endpoints Overview

#### Public Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| POST | `/api/v1/auth/register` | Register new user |
| POST | `/api/v1/auth/login` | Login user |
| GET | `/api/v1/products` | Get all products |
| GET | `/api/v1/products/:id` | Get product by ID |

#### Protected Endpoints (Require Authentication)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/users/profile` | Get current user profile |
| PUT | `/api/v1/users/profile` | Update current user profile |
| GET | `/api/v1/users` | Get all users |
| GET | `/api/v1/users/:id` | Get user by ID |
| DELETE | `/api/v1/users/:id` | Delete user |
| POST | `/api/v1/products` | Create product |
| GET | `/api/v1/products/my` | Get user's products |
| PUT | `/api/v1/products/:id` | Update product |
| DELETE | `/api/v1/products/:id` | Delete product |

### Example Requests

#### 1. Register a User

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john.doe@example.com",
    "username": "johndoe",
    "password": "password123",
    "first_name": "John",
    "last_name": "Doe"
  }'
```

#### 2. Login

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john.doe@example.com",
    "password": "password123"
  }'
```

**Response:**
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": { ... }
  }
}
```

#### 3. Create Product (with JWT)

```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop",
    "description": "High-performance laptop",
    "price": 1299.99,
    "stock": 50,
    "category": "Electronics",
    "sku": "LAP-001"
  }'
```

### Postman Collection

Import the Postman collection from [api/postman_collection.json](api/postman_collection.json) for a complete set of pre-configured requests.

For detailed examples, see [api/examples.md](api/examples.md)

---

## Task Automation

### Using Taskfile

```bash
# View all available tasks
task --list

# Install dependencies
task install

# Run in development mode
task dev

# Build the application
task build

# Run tests
task test

# Docker operations
task docker-up
task docker-down
task docker-logs
```

### Using Makefile

```bash
# View all available commands
make help

# Run in development mode
make dev

# Build the application
make build

# Run tests
make test

# Docker operations
make docker-up
make docker-down
```

---

## Deployment

### Docker Deployment

1. **Build and start containers**
   ```bash
   docker-compose up -d
   ```

2. **View logs**
   ```bash
   docker-compose logs -f
   ```

3. **Stop containers**
   ```bash
   docker-compose down
   ```

### Production Deployment Checklist

- [ ] Change `JWT_SECRET` to a strong random value
- [ ] Set `GIN_MODE=release`
- [ ] Enable SSL/TLS for database connections
- [ ] Set up proper logging and monitoring
- [ ] Configure reverse proxy (Nginx/Traefik)
- [ ] Set up CI/CD pipeline
- [ ] Enable database backups
- [ ] Configure rate limiting
- [ ] Set up health checks and alerts

---

## Testing

### Run Tests

```bash
# Run all tests
go test -v ./...

# Run tests with coverage
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

### Manual Testing

1. Start the server
2. Import Postman collection from `api/postman_collection.json`
3. Follow the request order: Register → Login → Use Token for other endpoints

---

## Database Schema

### Users Table

| Column | Type | Constraints |
|--------|------|-------------|
| id | SERIAL | PRIMARY KEY |
| email | VARCHAR | UNIQUE, NOT NULL |
| username | VARCHAR | UNIQUE, NOT NULL |
| password | VARCHAR | NOT NULL (hashed) |
| first_name | VARCHAR | |
| last_name | VARCHAR | |
| role | VARCHAR | DEFAULT 'user' |
| is_active | BOOLEAN | DEFAULT true |
| created_at | TIMESTAMP | |
| updated_at | TIMESTAMP | |
| deleted_at | TIMESTAMP | Soft delete |

### Products Table

| Column | Type | Constraints |
|--------|------|-------------|
| id | SERIAL | PRIMARY KEY |
| name | VARCHAR | NOT NULL |
| description | TEXT | |
| price | DECIMAL | NOT NULL |
| stock | INTEGER | DEFAULT 0 |
| category | VARCHAR | |
| sku | VARCHAR | UNIQUE |
| user_id | INTEGER | FOREIGN KEY → users(id) |
| created_at | TIMESTAMP | |
| updated_at | TIMESTAMP | |
| deleted_at | TIMESTAMP | Soft delete |

---

## Security Features

- **Password Hashing**: bcrypt with salt
- **JWT Authentication**: Secure token-based auth
- **CORS Protection**: Configurable origins
- **SQL Injection Prevention**: GORM parameterized queries
- **Input Validation**: Gin binding validation
- **Error Handling**: No sensitive data exposure
- **Soft Delete**: Data retention for auditing

---

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## Troubleshooting

### Database Connection Issues

```bash
# Check if PostgreSQL is running
docker ps | grep postgres

# View database logs
docker-compose logs postgres

# Connect to database
docker-compose exec postgres psql -U postgres -d go_rest_api
```

### Port Already in Use

```bash
# Find process using port 8080
lsof -i :8080  # macOS/Linux
netstat -ano | findstr :8080  # Windows

# Kill the process or change SERVER_PORT in .env
```

---

## Performance Considerations

- **Database Indexing**: Indexes on email, username, SKU for fast lookups
- **Connection Pooling**: GORM handles connection pooling
- **Pagination**: Implemented for list endpoints
- **Lazy Loading**: Relations loaded only when needed

---

## Roadmap

- [ ] Add comprehensive unit tests
- [ ] Implement rate limiting middleware
- [ ] Add Redis caching layer
- [ ] Implement refresh tokens
- [ ] Add email verification
- [ ] Integrate with cloud storage for files
- [ ] Add GraphQL support
- [ ] Implement WebSocket support
- [ ] Add metrics and monitoring (Prometheus)
- [ ] CI/CD pipeline with GitHub Actions

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 📫 Connect With Me

- 📱 Telegram: [@primeakash](https://t.me/primeakash)
- 💼 LinkedIn: [Nasir Hossain Akash](https://linkedin.com/in/nasirhossainakash)
- 🌐 Website: [BotsGalaxy.com](https://botsgalaxy.com)
- 📧 Email: admin@botsgalaxy.com

---

## Acknowledgments

- [Gin Web Framework](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [JWT-Go](https://github.com/golang-jwt/jwt)
- [PostgreSQL](https://www.postgresql.org/)

---

<div align="center">

**If you found this project helpful, please give it a ⭐**

Made with ❤️ using Go by [Nasir Hossain Akash](https://botsgalaxy.com)

</div>
