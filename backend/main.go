package main

import (
	"log"
	"os"

	"github.com/anton-fuji/skycast-app/backend/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".envファイルが読み込めません")
	}

	r := gin.Default()

	// CORSの設定
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}

	// エンドポイント
	api := r.Group("/api")
	{
		api.GET("/weather/:city", handler.GetWeatherByCity)
		api.GET("/health", handler.GetHealth)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
