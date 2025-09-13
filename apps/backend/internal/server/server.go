package server

// Imports
import (
	"atlance/weatherapp/internal/utils/fetch"
	"encoding/json"
	"log"
	"net/http"
)

// Variable Initializations
const (
	domain = "localhost"
)

// Weather response struct
type WeatherResponse struct {
	Location string  `json:"location"`
	TempC    string `json:"temp_c"`
	TempF    string `json:"temp_f"`
	Condition string `json:"condition"`
}

// Handlers

// Default handler
func index(res http.ResponseWriter, req *http.Request) {
	http.FileServer(http.Dir("../../frontend/dist")).ServeHTTP(res, req);
}

//API handler
func apiIndex(res http.ResponseWriter, req *http.Request) {
	w := WeatherResponse{}

	// Fetch weather data
	data := fetch.FetchWeatherData();

	// Populate weather response struct
	w.Location = "Farmington, MO"
	w.TempC = ""
	w.TempF = data
	w.Condition = ""

	// Convert to JSON
	err := json.NewEncoder(res).Encode(w);
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err);
	}
}


// Start creates a new server instance
func Start(port string) {

	// Root Handler
	http.HandleFunc("/", index);

	// Api Handler
	http.HandleFunc(domain + "/api", apiIndex);

	// Listen and Serve
	http.ListenAndServe(":" + port, nil);
}