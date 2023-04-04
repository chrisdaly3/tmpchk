package main

import (
	"fmt"
	"os"

	color "github.com/fatih/color"
	flag "github.com/ogier/pflag"
)

// Available flags for paramaters
var (
	location string
)

// Initializes available flags
func init() {
	flag.StringVarP(&location, "location", "l", "Philadelphia", "Search by City")
}
func main() {
	// Parses all flags defined in os.Args[1:], MUST be called before flags are accessed by program.
	flag.Parse()

	// If user does not supply the flags, print app use
	if flag.NFlag() == 0 {
		printUse()
	}

	coordLat, coordLon, err := getCoordinates(location)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		fmt.Println("Exited status 1")
	}

	response := getWeather(coordLat, coordLon)
	color.Blue(`Latitude: %v, Longitude: %v`, response.Latitude, response.Longitude)
	color.Green(`Timezone: %s`, response.Timezone)
	color.Magenta(`Current Weather - Temp: %v Fahrenheit`, response.CurrentWeather["temperature"])
	color.Magenta(`Current Weather - Windspeed: %v mph`, response.CurrentWeather["windspeed"])
}

func printUse() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options: ")
	flag.PrintDefaults()
	os.Exit(1)
}
