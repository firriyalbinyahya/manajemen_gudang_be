package repository

import (
	"manajemen_gudang_be/entity"

	"gorm.io/gorm"
)

type ReportRepository struct {
	DB *gorm.DB
}

func NewReportRepository(db *gorm.DB) *ReportRepository {
	return &ReportRepository{DB: db}
}

func (rr *ReportRepository) GetReportSummary() (entity.ReportSummary, error) {
	var summary entity.ReportSummary

	if err := rr.DB.Model(&entity.Product{}).Count(&summary.TotalProducts).Error; err != nil {
		return summary, err
	}

	if err := rr.DB.Model(&entity.Product{}).Select("SUM(quantity)").Scan(&summary.TotalStock).Error; err != nil {
		return summary, err
	}

	if err := rr.DB.Model(&entity.Product{}).Where("quantity <= ?", 10).Count(&summary.LowStockItems).Error; err != nil {
		return summary, err
	}

	return summary, nil
}
