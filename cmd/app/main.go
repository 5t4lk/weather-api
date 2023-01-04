package main

import (
	"awesomeProject4/internal/api/darksky"
	"awesomeProject4/internal/settings"
	"awesomeProject4/internal/user"
	"log"
)

func main() {
	fileData, err := settings.ScanFile("settings.txt")
	if err != nil {
		log.Fatalf("error while scanning file: %s", err)
	}

	latitude, longitude, apiKey, apiHost, err := settings.EjectData(fileData)
	if err != nil {
		log.Fatalf("error while ejecting data: %s", err)
	}

	weatherBytes, err := darksky.MakeRequest(latitude, longitude, apiKey, apiHost)
	if err != nil {
		log.Fatalf("error while receiving data from API: %s", err)
	}

	clearWeather, err := darksky.ClearJSON(weatherBytes)
	if err != nil {
		log.Fatalf("error while clearing JSON: %s", err)
	}

	if err = user.Output(clearWeather); err != nil {
		log.Fatalf("error while showing weather: %s", err)
	}
}
