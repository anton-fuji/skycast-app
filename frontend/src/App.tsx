import "./App.css";
import { Cloud } from "lucide-react";
import SearchForm from "../src/components/SerchForm"; // ←ファイル名がSerchFormならこのままでOK
import { useWeather } from "./hooks/useWeather"; // useWeatherのインポートを追加

function App() {
  const { loading, searchWeather } = useWeather();

  const handleSearch = (city: string) => {
    searchWeather(city); 
  };

  return (
    <div className="p-4 bg-gray-50 items-center flex flex-col min-h-screen justify-center">
      <div className="text-center mb-8">
        <div className="items-center space-x-4 flex justify-center mb-8">
          <Cloud className="h-8 w-8 text-blue-500" />
          <h1 className="text-3xl font-bold justify-center text-gray-800">Skycast</h1>
        </div>
        <p className="text-gray-600">今日の天気</p>
      </div>

      <div>
        <SearchForm onSerch={handleSearch} loading={loading} />
      </div>
    </div>
  );
}

export default App;
