package main

// Import necessary packages
import (
	"atlance/weatherapp/api/server"
	"os"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Set environment variables
	PORT := os.Getenv("PORT")

	// Start server
	server.Start(PORT);
}