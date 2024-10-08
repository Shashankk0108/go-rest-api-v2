package service

import (
	"errors"

	"github.com/botsgalaxy/go-rest-api-v2/internal/models"
	"github.com/botsgalaxy/go-rest-api-v2/internal/repository"
)

// ProductService handles product business logic
type ProductService interface {
	CreateProduct(req *models.CreateProductRequest, userID uint) (*models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	GetAllProducts(limit, offset int) ([]models.Product, error)
	GetProductsByCategory(category string, limit, offset int) ([]models.Product, error)
	GetUserProducts(userID uint, limit, offset int) ([]models.Product, error)
	UpdateProduct(id uint, req *models.UpdateProductRequest, userID uint) (*models.Product, error)
	DeleteProduct(id uint, userID uint) error
}

type productService struct {
	productRepo repository.ProductRepository
}

// NewProductService creates a new product service
func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

func (s *productService) CreateProduct(req *models.CreateProductRequest, userID uint) (*models.Product, error) {
	product := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
		SKU:         req.SKU,
		UserID:      userID,
	}

	if err := s.productRepo.Create(product); err != nil {
		return nil, errors.New("failed to create product")
	}

	return product, nil
}

func (s *productService) GetProductByID(id uint) (*models.Product, error) {
	return s.productRepo.FindByID(id)
}

func (s *productService) GetAllProducts(limit, offset int) ([]models.Product, error) {
	return s.productRepo.FindAll(limit, offset)
}

func (s *productService) GetProductsByCategory(category string, limit, offset int) ([]models.Product, error) {
	return s.productRepo.FindByCategory(category, limit, offset)
}

func (s *productService) GetUserProducts(userID uint, limit, offset int) ([]models.Product, error) {
	return s.productRepo.FindByUserID(userID, limit, offset)
}

func (s *productService) UpdateProduct(id uint, req *models.UpdateProductRequest, userID uint) (*models.Product, error) {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Check if user owns the product
	if product.UserID != userID {
		return nil, errors.New("unauthorized to update this product")
	}

	// Update fields if provided
	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Price > 0 {
		product.Price = req.Price
	}
	if req.Stock >= 0 {
		product.Stock = req.Stock
	}
	if req.Category != "" {
		product.Category = req.Category
	}

	if err := s.productRepo.Update(product); err != nil {
		return nil, errors.New("failed to update product")
	}

	return product, nil
}

func (s *productService) DeleteProduct(id uint, userID uint) error {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return err
	}

	// Check if user owns the product
	if product.UserID != userID {
		return errors.New("unauthorized to delete this product")
	}

	return s.productRepo.Delete(id)
}
