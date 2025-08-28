package routes

import (
	"manajemen_gudang_be/controller"
	"manajemen_gudang_be/middleware"
	"manajemen_gudang_be/repository"
	"manajemen_gudang_be/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProductRoutes(router *gin.RouterGroup, db *gorm.DB) {
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	product := router.Group("/products")
	product.Use(middleware.AuthUserMiddleware())
	{
		product.POST("", productController.CreateProductHandler)
		product.GET("", productController.GetAllProductsHandler)
		product.GET("/:id", productController.GetProductByIDHandler)
		product.PUT("/:id", productController.UpdateProductHandler)
		product.DELETE("/:id", productController.DeleteProductHandler)
	}
}
