package server

// Imports
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"atlance/weatherapp/api/utils/fetch"
)

// Server struct
type Server struct {
	Port	string
	Home	string
	Title	string
}

type WeatherData struct {
	Temperature		string	`json:"temperature"`
}

// Weather API Handler function
func apiHandler(w http.ResponseWriter, r *http.Request) {
	
	// Set response header and write JSON response
	w.Header().Set("Content-Type", "application/json");
	data := WeatherData {
		Temperature: fetch.FetchWeatherData(),
	}

	resp, err := json.Marshal(data);

	// Handle errors
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError);
		return;
	}

	// Write response
	w.Write(resp);
}

// Vite Handler function
func viteHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./frontend/dist/index.html");
}

// Start creates a new server instance
func Start(port string) {
	// New server instance
	s := Server {
		Port:	fmt.Sprintf(":%v", port),
	}

	// Start server
	
	http.HandleFunc("/api", apiHandler);
	http.HandleFunc("/", viteHandler);
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v", s.Port), nil));
}