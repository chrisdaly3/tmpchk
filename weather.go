package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	color "github.com/fatih/color"
)

// This is the base URL we need to call for all API requests to open-meteo.
// All parameters will append to the URL request accordingly.
const (
	apiURL = "https://api.open-meteo.com/v1/forecast"
	osmURL = "https://nominatim.openstreetmap.org/search/?q="
)

type WeatherResponse struct {
	Latitude       float64                `json:"latitude"`
	Longitude      float64                `json:"longitude"`
	Elevation      float64                `json:"elevation"`
	GenerationTime float64                `json:"generationtime_ms"`
	UTCOffset      int                    `json:"utc_offset_seconds"`
	Timezone       string                 `json:"timezone"`
	TimezoneAbbr   string                 `json:"timezone_abbreviation"`
	Hourly         map[string]interface{} `json:"hourly"`
	HourlyUnits    map[string]string      `json:"hourly_units"`
	Daily          map[string]interface{} `json:"daily"`
	DailyUnits     map[string]string      `json:"daily_units"`
	CurrentWeather map[string]interface{} `json:"current_weather"`
}

type NominatimResponse struct {
	Lat         string `json:"lat"`
	Long        string `json:"lon"`
	DisplayName string `json:"display_name"`
}

func getCoordinates(location string) (string, string, error) {
	// Send the location request to obtain coordinates for weather request
	resp, err := http.Get(osmURL + location + "&format=json&limit=1")
	if err != nil {
		return "", "", fmt.Errorf("error obtaining coordinates for %s", location)
	}
	defer resp.Body.Close()

	// read the response body and parse JSON result
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("error reading response of coordinates result")
	}

	// Unmarshal the JSON object data into the NominatimResponse struct
	var coordResponses []NominatimResponse
	err = json.Unmarshal([]byte(body), &coordResponses)
	if err != nil {
		return "", "", err
	}

	// Extract lat + lon values from response for injection to weather call
	if coordResponses[0].Lat == "" || coordResponses[0].Long == "" {
		return "", "", fmt.Errorf("no results found for %s", location)
	}
	color.White("Showing Results for %s", coordResponses[0].DisplayName)
	return coordResponses[0].Lat, coordResponses[0].Long, nil
}

func getWeather(lat, lon string) WeatherResponse {
	// Build the URL depending on main.go flag attributes
	res, err := http.Get(apiURL + "?latitude=" + lat + "&longitude=" + lon + "&hourly=temperature_2m&current_weather=true&temperature_unit=fahrenheit&windspeed_unit=mph&precipitation_unit=inch&timezone=America%2FNew_York")
	if err != nil {
		log.Fatalf("Error calling server: %s\n", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading JSON response: %s\n", err)
	}

	var weather WeatherResponse
	json.Unmarshal(body, &weather)

	return weather
}
