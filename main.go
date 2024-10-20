package main

import (
	"fmt"
	"log"

	"github.com/botsgalaxy/go-rest-api-v2/internal/config"
	"github.com/botsgalaxy/go-rest-api-v2/internal/database"
	"github.com/botsgalaxy/go-rest-api-v2/internal/handler"
	"github.com/botsgalaxy/go-rest-api-v2/internal/middleware"
	"github.com/botsgalaxy/go-rest-api-v2/internal/repository"
	"github.com/botsgalaxy/go-rest-api-v2/internal/service"
	"github.com/gin-gonic/gin"
)

// @title Go REST API
// @version 2.0
// @description A production-ready REST API built with Go, Gin, and PostgreSQL
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func Run() error {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Initialize database
	db, err := database.NewDatabase(&cfg.Database)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	// Run migrations
	if err := db.AutoMigrate(); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db.DB)
	productRepo := repository.NewProductRepository(db.DB)

	// Initialize services
	authService := service.NewAuthService(userRepo, cfg.JWT.Secret, cfg.JWT.Expiration)
	userService := service.NewUserService(userRepo)
	productService := service.NewProductService(productRepo)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	productHandler := handler.NewProductHandler(productService)

	// Set Gin mode
	gin.SetMode(cfg.Server.Mode)

	// Initialize router
	router := gin.New()

	// Global middleware
	router.Use(gin.Recovery())
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.CORSMiddleware(cfg.CORS.AllowedOrigins))

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Public routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// Public product routes
		products := v1.Group("/products")
		{
			products.GET("", productHandler.GetProducts)
			products.GET("/:id", productHandler.GetProductByID)
		}

		// Protected routes (require authentication)
		authenticated := v1.Group("")
		authenticated.Use(middleware.AuthMiddleware(cfg.JWT.Secret))
		{
			// User routes
			users := authenticated.Group("/users")
			{
				users.GET("/profile", userHandler.GetProfile)
				users.PUT("/profile", userHandler.UpdateProfile)
				users.GET("", userHandler.GetUsers)
				users.GET("/:id", userHandler.GetUserByID)
				users.DELETE("/:id", userHandler.DeleteUser)
			}

			// Protected product routes
			authenticatedProducts := authenticated.Group("/products")
			{
				authenticatedProducts.POST("", productHandler.CreateProduct)
				authenticatedProducts.GET("/my", productHandler.GetMyProducts)
				authenticatedProducts.PUT("/:id", productHandler.UpdateProduct)
				authenticatedProducts.DELETE("/:id", productHandler.DeleteProduct)
			}
		}
	}

	// Start server
	address := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Server starting on %s", address)

	return router.Run(address)
}

func main() {
	log.Println("🚀 Go REST API v2.0 - Starting...")

	if err := Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
