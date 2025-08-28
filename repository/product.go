package repository

import (
	"manajemen_gudang_be/entity"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (pr *ProductRepository) Create(product *entity.Product) error {
	return pr.DB.Create(product).Error
}

func (pr *ProductRepository) GetPaginatedAndFiltered(page, limit int, status, lowStock, search string) ([]entity.Product, int64, error) {
	var products []entity.Product
	var totalItems int64

	query := pr.DB.Model(&entity.Product{})

	if status != "" {
		query = query.Where("status = ?", status)
	}
	if lowStock == "true" {
		query = query.Where("quantity <= ?", 10)
	}
	if search != "" {
		keyword := "%" + search + "%"
		query = query.Where("product_name LIKE ? OR sku LIKE ? OR location LIKE ?", keyword, keyword, keyword)
	}

	if err := query.Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).Offset((page - 1) * limit).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, totalItems, nil
}

func (pr *ProductRepository) GetByID(id uint64) (*entity.Product, error) {
	var product entity.Product
	if err := pr.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (pr *ProductRepository) Update(product *entity.Product) error {
	return pr.DB.Save(product).Error
}

func (pr *ProductRepository) Delete(id uint64) error {
	return pr.DB.Delete(&entity.Product{}, id).Error
}
