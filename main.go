package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
		Pressure int     `json:"pressure"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
}

func main() {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENWEATHER_API_KEY belum diset")
	}

	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		city := r.URL.Query().Get("city")
		if city == "" {
			http.Error(w, "Parameter ?city= wajib diisi", http.StatusBadRequest)
			return
		}

		url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", city, apiKey)
		resp, err := http.Get(url)
		if err != nil || resp.StatusCode != 200 {
			http.Error(w, "Gagal mengambil data cuaca", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		var weather WeatherResponse
		if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
			http.Error(w, "Gagal decode data cuaca", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
			<html>
			<head><title>Cuaca di %s</title></head>
			<body style="font-family: Arial, sans-serif; margin: 40px;">
				<h1>Cuaca di %s</h1>
				<p><strong>Deskripsi:</strong> %s</p>
				<p><strong>Suhu:</strong> %.2f °C</p>
				<p><strong>Kelembapan:</strong> %d%%</p>
				<p><strong>Tekanan:</strong> %d hPa</p>
				<p><strong>Kecepatan Angin:</strong> %.2f m/s</p>
				<p><strong>Cuaca:</strong> %s</p>
			</body>
			</html>
		`, weather.Name, weather.Name, weather.Weather[0].Description, weather.Main.Temp, weather.Main.Humidity, weather.Main.Pressure, weather.Wind.Speed, weather.Weather[0].Main)
	})

	http.HandleFunc("/weather/json", func(w http.ResponseWriter, r *http.Request) {
		city := r.URL.Query().Get("city")
		if city == "" {
			http.Error(w, "Parameter ?city= wajib diisi", http.StatusBadRequest)
			return
		}

		url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", city, apiKey)
		resp, err := http.Get(url)
		if err != nil || resp.StatusCode != 200 {
			http.Error(w, "Gagal mengambil data cuaca", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		var weather WeatherResponse
		if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
			http.Error(w, "Gagal decode data cuaca", http.StatusInternalServerError)
			return
		}

		result := map[string]interface{}{
			"city":        weather.Name,
			"temperature": weather.Main.Temp,
			"humidity":    weather.Main.Humidity,
			"pressure":    weather.Main.Pressure,
			"wind_speed":  weather.Wind.Speed,
			"weather":     weather.Weather[0].Main,
			"description": weather.Weather[0].Description,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
