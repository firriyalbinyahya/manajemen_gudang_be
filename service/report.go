package service

import (
	"manajemen_gudang_be/entity"
	"manajemen_gudang_be/repository"
)

type ReportService struct {
	ReportRepository *repository.ReportRepository
}

func NewReportService(reportRepository *repository.ReportRepository) *ReportService {
	return &ReportService{ReportRepository: reportRepository}
}

func (rs *ReportService) GetReportSummary() (entity.ReportSummary, error) {
	return rs.ReportRepository.GetReportSummary()
}
