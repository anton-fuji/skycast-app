// tailwind.config.js
/** @type {import('tailwindcss').Config} */
export default {
	content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
	theme: {
		extend: {
			colors: {
				"sunny-start": "#74b9ff",
				"sunny-end": "#0984e3",
				"rainy-start": "#636e72",
				"rainy-end": "#2d3436",
				"cloudy-start": "#a29bfe",
				"cloudy-end": "#6c5ce7",
			},
			backgroundImage: {
				"sunny-gradient": "linear-gradient(135deg, #74b9ff 0%, #0984e3 100%)",
				"rainy-gradient": "linear-gradient(135deg, #636e72 0%, #2d3436 100%)",
				"cloudy-gradient": "linear-gradient(135deg, #a29bfe 0%, #6c5ce7 100%)",
			},
		},
	},
	plugins: [],
};
