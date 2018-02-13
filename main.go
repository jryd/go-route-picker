package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"googlemaps.github.io/maps"
)

// googleMapsAPIKey is a global variable to store the API key we need to use
// the Google maps API
var googleMapsAPIKey string

// startAddress is a global variable to store our origin/starting address
var startAddress string

// endAddress is a global variable to store our desitnation/end address
var endAddress string

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	googleMapsAPIKey = os.Getenv("GOOGLE_MAPS_API_KEY")
	startAddress = os.Getenv("START_ADDRESS")
	endAddress = os.Getenv("END_ADDRESS")

	if googleMapsAPIKey == "" || startAddress == "" || endAddress == "" {
		log.Fatal("One or more environment variables have not been set. Please see the .env.example file for fields required.")
	}
}

func main() {

	c, err := maps.NewClient(maps.WithAPIKey(googleMapsAPIKey))

	if err != nil {
		log.Fatalf("Error creating new map client: %s", err)
	}

	r := &maps.DirectionsRequest{
		Origin:      startAddress,
		Destination: endAddress,
	}

	route, _, err := c.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("Error getting directions: %s", err)
	}

	if strings.Contains(route[0].Legs[0].Steps[1].HTMLInstructions, "right") {
		fmt.Println("Go the regular way! ETA to home:", route[0].Legs[0].Duration)
	} else {
		fmt.Println("Go the alternative way! ETA to home:", route[0].Legs[0].Duration)
	}
}
