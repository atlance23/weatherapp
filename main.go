package main

// Import necessary packages
import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"os"
)

// Variables
const url = "https://api.weather.gov/stations/KFAM/observations/latest";

// Data struct
type Response struct {
	Properties struct {
		Temperature struct {
			Value float64 `json:"value"`;	
		}
		WindSpeed struct {
			Value float64 `json:"value"`;	
		}
		RelativeHumidity struct {
			Value float64 `json:"value"`;	
		}
	}
}

// Format string function
func formatResult(data float64) string {
	return fmt.Sprintf("%v", data);
}

// Convert to Farenheit function
func toFarenheit(celsius float64) float64 {
	return (celsius * 9/5) + 32;
}

// Fecth weather data function
func fetchWeatherData() string {

	// Define variables
	var r Response;

	// Make the HTTP GET request
	resp, err := http.Get(url);

	// Handle errors
	if err != nil {
		fmt.Println("Error sending GET request:", err);
		os.Exit(1);
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body);

	// Handle errors
	if err != nil {
		fmt.Println("Error reading response body:", err);
		os.Exit(1);
	}

	// Unmarshal the JSON response into the struct
	err = json.Unmarshal(body, &r);

	// Handle errors
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err);
		os.Exit(1);
	}

	// Convert to Farenheit
	temp := toFarenheit(r.Properties.Temperature.Value);

	// Ensure the response body is closed after reading
	defer resp.Body.Close();

	// Return fetched data
	return formatResult(temp);
}

func main() {
	// Variables
	var userInput string;
	var fetchedData string;

	// Print welcome message
	fmt.Println("Weather App");
	fmt.Println("Would you like to fetch the latest weather data? (y/n):");
	fmt.Print(">>> ");

	// Get user input
	fmt.Scanln(&userInput);

	// If user input is 'y', fetch weather data
	switch userInput {
		case "y":
			// Display fetched data
			fetchedData = fmt.Sprintf("Temperature in Farmington, MO is: %v °F", fetchWeatherData());
		case "Y":
			// Display fetched data
			fetchedData = fmt.Sprintf("Temperature in Farmington, MO is: %v °F", fetchWeatherData());
		case "N":
			fmt.Println("Exiting program.");
			os.Exit(0);
		case "n":
			fmt.Println("Exiting program.");
			os.Exit(0);
		default:
			fmt.Println("Invalid input. Please enter 'y' or 'n'.");
			os.Exit(1);
	}

	// Print fetched data
	fmt.Println(fetchedData);
}