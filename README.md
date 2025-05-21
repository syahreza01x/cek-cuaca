# Weather API Go

Aplikasi backend sederhana untuk menampilkan data cuaca berdasarkan nama kota menggunakan Go dan OpenWeather API.

## Fitur

- Ambil data cuaca real-time dari OpenWeather API.
- Tampil dalam format HTML yang rapi.
- Bisa juga mendapatkan data dalam format JSON.
- Mudah digunakan hanya dengan memasukkan nama kota lewat URL parameter.

## Cara Pakai

1. Pastikan sudah punya API Key dari [OpenWeather](https://openweathermap.org/api).
2. Set environment variable `OPENWEATHER_API_KEY` dengan API Key kamu:
   ```bash
   export OPENWEATHER_API_KEY=apikey_anda
3. Jalankan Program 
    - go run main.go
4. Buka Browser
    - Untuk tampilan web : http://localhost:8000/weather?city=KotaAnda
    - Untuk tampilan JSON : http://localhost:8000/weather/json?city=KotaAnda