package controller

import (
	"errors"
	"manajemen_gudang_be/service"
	"manajemen_gudang_be/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportController struct {
	ReportService *service.ReportService
}

func NewReportController(reportService *service.ReportService) *ReportController {
	return &ReportController{ReportService: reportService}
}

func (rc *ReportController) GetReportSummary(c *gin.Context) {
	summary, err := rc.ReportService.GetReportSummary()
	if err != nil {
		response.BuildErrorResponse(c, errors.New("failed to get dashboard summary"))
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Success to get report summary", summary, nil)
}
