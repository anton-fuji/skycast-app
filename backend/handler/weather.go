package handler

import (
	"net/http"
	"strings"

	"github.com/anton-fuji/skycast-app/backend/models"
	"github.com/anton-fuji/skycast-app/backend/services"
	"github.com/gin-gonic/gin"
)

// APIの動作確認
func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"message": "Skycast is running",
	})
}

// 都市ごとのデータ
func GetWeatherByCity(c *gin.Context) {
	rowCity := strings.ToLower(c.Param("city"))

	city, ok := models.CityAlias[rowCity]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "都市が選択されていません",
		})
		return
	}

	WeatherResponse, statusCode, err := services.GetWeather(city)
	if err != nil{
		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK,  WeatherResponse)

}