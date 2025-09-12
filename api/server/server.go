package server

// Imports
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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

// Root Handler function
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! The time is %s", time.Now().Format(time.RFC1123));
}

// Weather API Handler function
func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Weather data object
	wd := WeatherData {
		Temperature: fetch.FetchWeatherData(),
	}
	
	// JSON encode the response
	encoder := json.NewEncoder(w);

	// Encode the weather data
	err := encoder.Encode(wd);

	// Error handling
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}
}

// Start creates a new server instance
func Start() {
	// New server instance
	s := Server {
		Port:	":80",
		Home:	"/",
		Title:	"Weather API Server",
	}

	// Start server
	http.HandleFunc("/", handler);
	http.HandleFunc("/api", apiHandler);
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v", s.Port), nil));
}