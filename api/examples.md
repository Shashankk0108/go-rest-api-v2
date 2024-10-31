# API Examples

This document provides examples of how to use the Go REST API.

## Base URL

```
http://localhost:8080/api/v1
```

## Authentication

All protected endpoints require a JWT token in the Authorization header:

```
Authorization: Bearer <your_token>
```

---

## Health Check

Check if the API is running:

```bash
curl http://localhost:8080/health
```

**Response:**
```json
{
  "status": "ok",
  "message": "Server is running"
}
```

---

## Authentication Endpoints

### Register a New User

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

**Response:**
```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "id": 1,
    "email": "john.doe@example.com",
    "username": "johndoe",
    "first_name": "John",
    "last_name": "Doe",
    "role": "user",
    "is_active": true,
    "created_at": "2025-11-20T00:00:00Z"
  }
}
```

### Login

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
    "user": {
      "id": 1,
      "email": "john.doe@example.com",
      "username": "johndoe",
      "first_name": "John",
      "last_name": "Doe",
      "role": "user",
      "is_active": true,
      "created_at": "2025-11-20T00:00:00Z"
    }
  }
}
```

---

## User Endpoints

### Get Current User Profile

```bash
curl http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer <your_token>"
```

### Update Profile

```bash
curl -X PUT http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Smith",
    "username": "johnsmith"
  }'
```

### Get All Users (with pagination)

```bash
curl "http://localhost:8080/api/v1/users?limit=10&offset=0" \
  -H "Authorization: Bearer <your_token>"
```

### Get User by ID

```bash
curl http://localhost:8080/api/v1/users/1 \
  -H "Authorization: Bearer <your_token>"
```

### Delete User

```bash
curl -X DELETE http://localhost:8080/api/v1/users/1 \
  -H "Authorization: Bearer <your_token>"
```

---

## Product Endpoints

### Create Product (Protected)

```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop",
    "description": "High-performance laptop for developers",
    "price": 1299.99,
    "stock": 50,
    "category": "Electronics",
    "sku": "LAP-001"
  }'
```

### Get All Products (Public)

```bash
curl "http://localhost:8080/api/v1/products?limit=10&offset=0"
```

### Get Products by Category (Public)

```bash
curl "http://localhost:8080/api/v1/products?category=Electronics&limit=10"
```

### Get Product by ID (Public)

```bash
curl http://localhost:8080/api/v1/products/1
```

### Get My Products (Protected)

```bash
curl "http://localhost:8080/api/v1/products/my?limit=10&offset=0" \
  -H "Authorization: Bearer <your_token>"
```

### Update Product (Protected)

```bash
curl -X PUT http://localhost:8080/api/v1/products/1 \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Gaming Laptop",
    "description": "High-performance gaming laptop",
    "price": 1499.99,
    "stock": 30
  }'
```

### Delete Product (Protected)

```bash
curl -X DELETE http://localhost:8080/api/v1/products/1 \
  -H "Authorization: Bearer <your_token>"
```

---

## Error Responses

All error responses follow this format:

```json
{
  "success": false,
  "error": "Error message here"
}
```

### Common HTTP Status Codes

- `200 OK` - Request successful
- `201 Created` - Resource created successfully
- `400 Bad Request` - Invalid request data
- `401 Unauthorized` - Missing or invalid authentication
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server error

---

## Using with JavaScript/Fetch

```javascript
// Register
const register = async () => {
  const response = await fetch('http://localhost:8080/api/v1/auth/register', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      email: 'john.doe@example.com',
      username: 'johndoe',
      password: 'password123',
      first_name: 'John',
      last_name: 'Doe'
    })
  });
  return await response.json();
};

// Login and get token
const login = async () => {
  const response = await fetch('http://localhost:8080/api/v1/auth/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      email: 'john.doe@example.com',
      password: 'password123'
    })
  });
  const data = await response.json();
  return data.data.token;
};

// Use token for authenticated requests
const getProfile = async (token) => {
  const response = await fetch('http://localhost:8080/api/v1/users/profile', {
    headers: {
      'Authorization': `Bearer ${token}`
    }
  });
  return await response.json();
};
```
