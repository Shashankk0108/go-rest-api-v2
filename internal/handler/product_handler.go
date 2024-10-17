package handler

import (
	"strconv"

	"github.com/botsgalaxy/go-rest-api-v2/internal/middleware"
	"github.com/botsgalaxy/go-rest-api-v2/internal/models"
	"github.com/botsgalaxy/go-rest-api-v2/internal/service"
	"github.com/botsgalaxy/go-rest-api-v2/internal/utils"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.CreateProductRequest true "Product details"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req models.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	product, err := h.productService.CreateProduct(&req, userID)
	if err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	utils.SuccessResponse(c, 201, "Product created successfully", product)
}

// GetProducts godoc
// @Summary Get all products
// @Description Get list of all products (with pagination)
// @Tags products
// @Accept json
// @Produce json
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Param category query string false "Category"
// @Success 200 {object} utils.Response
// @Router /products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	category := c.Query("category")

	var products []models.Product
	var err error

	if category != "" {
		products, err = h.productService.GetProductsByCategory(category, limit, offset)
	} else {
		products, err = h.productService.GetAllProducts(limit, offset)
	}

	if err != nil {
		utils.ErrorResponse(c, 500, err.Error())
		return
	}

	utils.SuccessResponse(c, 200, "Products retrieved successfully", products)
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Get product details by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /products/{id} [get]
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, 400, "Invalid product ID")
		return
	}

	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, 404, err.Error())
		return
	}

	utils.SuccessResponse(c, 200, "Product retrieved successfully", product)
}

// GetMyProducts godoc
// @Summary Get user's products
// @Description Get all products created by the authenticated user
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /products/my [get]
func (h *ProductHandler) GetMyProducts(c *gin.Context) {
	userID := middleware.GetUserID(c)
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	products, err := h.productService.GetUserProducts(userID, limit, offset)
	if err != nil {
		utils.ErrorResponse(c, 500, err.Error())
		return
	}

	utils.SuccessResponse(c, 200, "Products retrieved successfully", products)
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update product details
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Param request body models.UpdateProductRequest true "Update details"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	userID := middleware.GetUserID(c)

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, 400, "Invalid product ID")
		return
	}

	var req models.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	product, err := h.productService.UpdateProduct(uint(id), &req, userID)
	if err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	utils.SuccessResponse(c, 200, "Product updated successfully", product)
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Delete a product
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	userID := middleware.GetUserID(c)

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, 400, "Invalid product ID")
		return
	}

	if err := h.productService.DeleteProduct(uint(id), userID); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	utils.SuccessResponse(c, 200, "Product deleted successfully", nil)
}
