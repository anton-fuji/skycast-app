import { useState } from "react";

interface WeatherData {
  location: string;
  temperature: number;
  description: string;
  humidity: number;
  windSpeed: number;
}

interface WeatherState {
  data: WeatherData | null;
  loading: boolean;
  error: string | null;
}

const fetchWeatherByCity = async (city: string): Promise<WeatherData> => {
  const response = await fetch(`http://localhost:8000/api/weather/${encodeURIComponent(city)}`);
  
  if (!response.ok) {
    throw new Error('天気情報の取得に失敗しました');
  }
  
  return response.json();
};

export const useWeather = () => {
  const [state, setState] = useState<WeatherState>({
    data: null,
    loading: false,
    error: null,
  });

  const searchWeather = async (city: string) => {
    setState(prev => ({ ...prev, loading: true, error: null }));

    try {
      const data = await fetchWeatherByCity(city);
      setState({ data, loading: false, error: null });
    } catch (error) {
      setState({
        data: null,
        loading: false,
        error: error instanceof Error ? error.message : 'An unexpected error occurred',
      });
    }
  };

  return {
    ...state,
    searchWeather,
  };
};