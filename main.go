package main

import (
	"fmt"
	"manajemen_gudang_be/config"
	"manajemen_gudang_be/config/database"
	"manajemen_gudang_be/entity"
	"manajemen_gudang_be/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func InitializeApp() *gin.Engine {
	config.InitConfig()

	r := gin.Default()

	db := database.ConnectDatabase()
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	routes.SetupRoutes(r, db)

	return r
}

func main() {
	app := InitializeApp()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000" // Default port
	}
	fmt.Println("Server is running on port " + port)
	app.Run(":" + port)
}
