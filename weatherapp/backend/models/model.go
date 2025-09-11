package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Data struct
type Data struct {
	Properties struct {
		Temperature struct {
			Value float64 `json:"value"`
		} `json:"temperature"`
	} `json:"properties"`
}


// Fetch weather method
func fetchWeatherData() ([]byte, error) {
	// Fetch the data from the API
	resp, err := http.Get("https://api.weather.gov/stations/KFAM/observations/latest")

	// Error handling
	if err != nil {
		return nil, fmt.Errorf("error fetching data from API: %s", err);
	}

	// Defer closing the response body
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	// Read and return the response body
	return io.ReadAll(resp.Body)
}

// Convert temperature from Celcius to Fahrenheit
func cToF(celsius float64) float64 {
	return ((float64(celsius) * 9.0)/ 5.0) + 32.0
}

// Set weather method
func (d *Data) setWeather(jsonData []byte) error{
	
	// Unmarshal JSON data into a map
	err := json.Unmarshal(jsonData, &d);

	// Handle errors
	if err != nil {
		fmt.Println("Error unmarshalling JSON to map:", err)
	}

	return err
}

// Main method
func main() {
	// Get current time
	currentTime := time.Now()
	utcTime := currentTime.UTC()
	currTime := fmt.Sprintf("Current UTC time: %s", utcTime)

	// Fetch weather data
	jsonData, err := fetchWeatherData()

	// Handle errors
	if err != nil {
		fmt.Println("Error calling fetch function", err)
		return
	}

	// Create a Properties instance
	var d Data

	// Unmarshal JSON data into Weather struct
	err = d.setWeather(jsonData)

	// Handle errors
	if err != nil {
		fmt.Println("Error setting weather data:", err)
		return
	}

	// Print the weather data
	fmt.Printf("The weather in Farmington, MO was: Temperature: %v °F, at %s\n", cToF(d.Properties.Temperature.Value), currTime)
}
