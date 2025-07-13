package services

import (
	"net/http"
	"os"
	"fmt"
	"encoding/json"

	"github.com/anton-fuji/skycast-app/backend/models"
)

func GetWeather(city string) (*models.WeatherResponse, int, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		return nil, http.StatusInternalServerError, fmt.Errorf("APIキーが設定されていません")
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&lang=ja&units=metric", city, apiKey)
	
	// HTTPリクエストの作成
	resp, err := http.Get(url)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("データの取得に失敗しました")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, http.StatusNotFound, fmt.Errorf("指定された都市が見つかりません")
	}

	var openWeatherMapData models.OpenWeatherMap
	if err := json.NewDecoder(resp.Body).Decode(&openWeatherMapData); err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("データの解析に失敗しました")
	}
	
	// Transform data
	response := &models.WeatherResponse{
		City:        openWeatherMapData.Name,
		Country:     openWeatherMapData.Country,
		Temperature: openWeatherMapData.Main.Temp,
		Description: openWeatherMapData.Weather[0].Description,
		Humidity:    openWeatherMapData.Main.Humidity,
		// Pressure:    openWeatherMapData.Main.Pressure,
		// WindSpeed:   openWeatherMapData.Wind.Speed,
		FeelsLike:   openWeatherMapData.Main.FeelsLike,
	}
	
	return response, http.StatusOK, nil
}
