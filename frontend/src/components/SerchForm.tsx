import React, { useState } from "react";
import { Search } from "lucide-react";

type SearchProps = {
  onSerch: (city: string) => void;
  loading: boolean;
};

const SearchForm: React.FC<SearchProps> = ({ onSerch, loading }) => {
  const [city, setCity] = useState("");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (city.trim()) {
      onSerch(city.trim());
    }
  };

  return (
    <form onSubmit={handleSubmit} className="w-full max-w-md">
      <div className="flex items-center border border-gray-300 rounded-lg overflow-hidden">
        <input
          type="text"
          value={city}
          onChange={(e) => setCity(e.target.value)}
          placeholder="地域を入力してください"
          className="w-full px-4 py-3 focus:outline-none"
          disabled={loading}
        />
        <button
          type="submit"
          disabled={loading || !city.trim()}
          className="p-3 text-gray-500 disabled:opacity-50"
        >
          <Search className="h-5 w-5" />
        </button>
      </div>
    </form>
  );
};

export default SearchForm;
