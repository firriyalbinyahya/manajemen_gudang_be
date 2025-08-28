package controller

import (
	"manajemen_gudang_be/entity"
	"manajemen_gudang_be/service"
	"manajemen_gudang_be/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (ac *AuthController) Register(c *gin.Context) {
	var req entity.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	err := ac.AuthService.Register(&req)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusCreated, "registration successful", nil, nil)
}

func (ac *AuthController) Login(c *gin.Context) {
	var request entity.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	result, err := ac.AuthService.Login(&request)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Login successful", result, nil)

}
