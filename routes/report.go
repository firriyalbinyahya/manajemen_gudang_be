package routes

import (
	"manajemen_gudang_be/controller"
	"manajemen_gudang_be/middleware"
	"manajemen_gudang_be/repository"
	"manajemen_gudang_be/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupReportRoutes(router *gin.RouterGroup, db *gorm.DB) {
	reportRepository := repository.NewReportRepository(db)
	reportService := service.NewReportService(reportRepository)
	reportController := controller.NewReportController(reportService)

	report := router.Group("/report")
	{
		report.GET("", middleware.AuthUserMiddleware(), reportController.GetReportSummary)
	}
}
