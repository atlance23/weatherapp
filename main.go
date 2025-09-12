package main

// Import necessary packages
import (
	"fmt"
	"os"
	"atlance/weatherapp/api/server"
	"atlance/weatherapp/api/utils/fetch"
)

func main() {
	// Variables
	var userInput string;
	var fetchedData string;

	// Start server
	go server.Start();

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
			fetchedData = fmt.Sprintf("Temperature in Farmington, MO is: %v °F", fetch.FetchWeatherData());
		case "Y":
			// Display fetched data
			fetchedData = fmt.Sprintf("Temperature in Farmington, MO is: %v °F", fetch.FetchWeatherData());
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