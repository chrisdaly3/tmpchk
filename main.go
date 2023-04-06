package main

import (
	"fmt"
	"os"
	"time"

	color "github.com/fatih/color"
	flag "github.com/ogier/pflag"
)

// Available flags for paramaters
var (
	location string
)

// Initializes available flags
func init() {
	flag.StringVarP(&location, "location", "l", "[nameOfCity,XY,USA]", "Search by location")
}
func main() {
	// Parses all flags defined in os.Args[1:], MUST be called before flags are accessed by program.
	flag.Parse()

	// If user does not supply the flags, print app use
	if flag.NFlag() == 0 {
		printUse()
	}

	// Capture compute time, and obtain coordinates of the location declared in --location using the Nominatim Geocoding API
	start := time.Now()
	coordLat, coordLon, err := getCoordinates(location)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		fmt.Println("Exited status 1")
	}

	// Obtain weather JSON object from Open-Meteo using coordinates from getCoordinates() and print to Console
	response := getWeather(coordLat, coordLon)
	end := time.Now()
	computeTime := end.Sub(start)
	color.Blue(`Latitude: %v, Longitude: %v`, response.Latitude, response.Longitude)
	color.Green(`Timezone: %s`, response.Timezone)
	color.Magenta(`Current Weather - Temp: %v Fahrenheit`, response.CurrentWeather["temperature"])
	color.Magenta(`Current Weather - Windspeed: %v mph`, response.CurrentWeather["windspeed"])
	color.White("Time to compute: %v", computeTime)
}

// function printUse() prints the default options available to the console
func printUse() {
	fmt.Printf("Usage: %s [-l]\n", os.Args[0])
	fmt.Println("Options: ")
	flag.PrintDefaults()
	os.Exit(1)
}
