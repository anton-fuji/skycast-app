package models

type WeatherResponse struct {
	City        string  `json:"city"`
	Country     string  `json:"country"`
	Temperature float64 `json:"temperature"`
	Description string  `json:"description"`
	Humidity    int     `json:"humidity"`
	// Pressure    int     `json:"pressure"`
	// WindSpeed   float64 `json:"windSpeed"`
	FeelsLike   float64 `json:"feelsLike"`
}

// OpenWeather APIからのレスポンス用の構造体
type OpenWeatherMap struct{
	Name string `json:"name"`
	Country string `json:"country"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity  int     `json:"humidity"`
		Pressure  int     `json:"pressure"`
	} `json:"main"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
}