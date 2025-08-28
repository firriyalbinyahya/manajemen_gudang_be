package controller

import (
	"errors"
	"manajemen_gudang_be/entity"
	"manajemen_gudang_be/service"
	"manajemen_gudang_be/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService *service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

func (pc *ProductController) CreateProductHandler(c *gin.Context) {
	var req entity.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BuildErrorResponse(c, errors.New("invalid request body"))
		return
	}

	if err := pc.ProductService.CreateProduct(&req); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}
	response.BuildSuccessResponse(c, http.StatusCreated, "Product created successfully", nil, nil)
}

func (pc *ProductController) GetAllProductsHandler(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	status := c.Query("status")
	lowStock := c.Query("low_stock")
	search := c.Query("search")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		response.BuildErrorResponse(c, errors.New("limit harus angka positif"))
		return
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		response.BuildErrorResponse(c, errors.New("page harus angka positif"))
		return
	}

	products, err := pc.ProductService.GetPaginatedProducts(page, limit, status, lowStock, search)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Success to get data of products", products, nil)
}

func (pc *ProductController) GetProductByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BuildErrorResponse(c, errors.New("id tidak valid"))
		return
	}
	product, err := pc.ProductService.GetProductByID(id)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Success to get data of product", product, nil)
}

func (pc *ProductController) UpdateProductHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BuildErrorResponse(c, errors.New("id tidak valid"))
		return
	}
	var req entity.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BuildErrorResponse(c, errors.New("invalid request body"))
		return
	}

	if err := pc.ProductService.UpdateProduct(id, &req); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}
	response.BuildSuccessResponse(c, http.StatusOK, "Product updated successfully", nil, nil)
}

func (pc *ProductController) DeleteProductHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BuildErrorResponse(c, errors.New("id tidak valid"))
		return
	}
	if err := pc.ProductService.DeleteProduct(id); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}
	response.BuildSuccessResponse(c, http.StatusOK, "Product deleted successfully", nil, nil)
}
