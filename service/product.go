package service

import (
	"manajemen_gudang_be/entity"
	"manajemen_gudang_be/repository"
	"math"
)

type ProductService struct {
	ProductRepository *repository.ProductRepository
}

func NewProductService(productRepository *repository.ProductRepository) *ProductService {
	return &ProductService{ProductRepository: productRepository}
}

func (ps *ProductService) CreateProduct(req *entity.CreateProductRequest) error {
	product := &entity.Product{
		ProductName: req.ProductName,
		SKU:         req.SKU,
		Quantity:    req.Quantity,
		Location:    req.Location,
		Status:      req.Status,
	}
	return ps.ProductRepository.Create(product)
}

func (ps *ProductService) GetPaginatedProducts(page, limit int, status, lowStock, search string) (*entity.PaginatedProduct, error) {
	products, totalItems, err := ps.ProductRepository.GetPaginatedAndFiltered(page, limit, status, lowStock, search)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	response := &entity.PaginatedProduct{
		Items:       products,
		CurrentPage: page,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
	}

	return response, nil
}

func (ps *ProductService) GetProductByID(id uint64) (*entity.Product, error) {
	return ps.ProductRepository.GetByID(id)
}

func (ps *ProductService) UpdateProduct(id uint64, req *entity.UpdateProductRequest) error {
	product, err := ps.ProductRepository.GetByID(id)
	if err != nil {
		return err
	}

	if req.ProductName != nil {
		product.ProductName = *req.ProductName
	}
	if req.Quantity != nil {
		product.Quantity = *req.Quantity
	}
	if req.Location != nil {
		product.Location = *req.Location
	}
	if req.Status != nil {
		product.Status = *req.Status
	}

	return ps.ProductRepository.Update(product)
}

func (ps *ProductService) DeleteProduct(id uint64) error {
	return ps.ProductRepository.Delete(id)
}
