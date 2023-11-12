package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv" 
	"github.com/tidwall/gjson"
)

const openWeatherMapURL = "http://api.openweathermap.org/data/2.5/weather"

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Check if the user provided a city as a command-line argument.
	if len(os.Args) < 2 {
		fmt.Println("Usage: weather <city>")
		os.Exit(1)
	}

	// Retrieve the city from the command-line arguments.
	city := os.Args[1]

	// Retrieve API key from environment variable
	apiKey := os.Getenv("API_KEY")

	// Call the getWeather function to retrieve weather data.
	weatherData, err := getWeather(city, apiKey)
	if err != nil {
		log.Fatal(err)
	}

	// Print the weather information.
	fmt.Printf("Weather in %s:\n", city)
	fmt.Printf("Description: %s\n", weatherData.Description)
	fmt.Printf("Temperature: %.2f°C\n", weatherData.Temperature)
	fmt.Printf("Humidity: %d%%\n", weatherData.Humidity)
	fmt.Printf("Wind Speed: %.2f m/s\n", weatherData.WindSpeed)
}

// WeatherData is a struct to represent the weather information.
type WeatherData struct {
	Description string
	Temperature float64
	Humidity    int
	WindSpeed   float64
}

// getWeather function retrieves weather data from the OpenWeatherMap API.
func getWeather(city, apiKey string) (WeatherData, error) {
	// Build the URL for the OpenWeatherMap API request.
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", openWeatherMapURL, city, apiKey)

	// Make an HTTP GET request to the OpenWeatherMap API.
	resp, err := http.Get(url)
	if err != nil {
		return WeatherData{}, err
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return WeatherData{}, err
	}

	// Check if the API request was successful (status code 200 OK).
	if resp.StatusCode != http.StatusOK {
		return WeatherData{}, fmt.Errorf("Failed to fetch weather data. Status code: %d", resp.StatusCode)
	}

	// Extract weather information from the JSON response using gjson.
	description := gjson.GetBytes(body, "weather.0.description").String()
	temperature := gjson.GetBytes(body, "main.temp").Float()
	humidity := gjson.GetBytes(body, "main.humidity").Int()
	windSpeed := gjson.GetBytes(body, "wind.speed").Float()

	// Create and return a WeatherData struct with the extracted information.
	return WeatherData{
		Description: description,
		Temperature: temperature,
		Humidity:    int(humidity),
		WindSpeed:   windSpeed,
	}, nil
}
