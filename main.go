package main

import (
	"fmt"
	"manajemen_gudang_be/config"
	"manajemen_gudang_be/config/database"
	"manajemen_gudang_be/entity"
	"manajemen_gudang_be/routes"
	"os"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func InitializeApp() *gin.Engine {
	config.InitConfig()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://simpel-warehouse-firriyal.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

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
