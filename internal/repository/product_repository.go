package repository

import (
	"errors"

	"github.com/botsgalaxy/go-rest-api-v2/internal/models"
	"gorm.io/gorm"
)

// ProductRepository handles product data access
type ProductRepository interface {
	Create(product *models.Product) error
	FindByID(id uint) (*models.Product, error)
	FindAll(limit, offset int) ([]models.Product, error)
	FindByUserID(userID uint, limit, offset int) ([]models.Product, error)
	FindByCategory(category string, limit, offset int) ([]models.Product, error)
	Update(product *models.Product) error
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new product repository
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("User").First(&product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) FindAll(limit, offset int) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("User").Limit(limit).Offset(offset).Find(&products).Error
	return products, err
}

func (r *productRepository) FindByUserID(userID uint, limit, offset int) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Where("user_id = ?", userID).Limit(limit).Offset(offset).Find(&products).Error
	return products, err
}

func (r *productRepository) FindByCategory(category string, limit, offset int) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Where("category = ?", category).Preload("User").Limit(limit).Offset(offset).Find(&products).Error
	return products, err
}

func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
