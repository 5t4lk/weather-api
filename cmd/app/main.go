package main

import (
	"awesomeProject4/internal/api/darksky"
	"awesomeProject4/internal/settings"
	"awesomeProject4/internal/user"
	"log"
)

func main() {
	if err := settings.Eject(); err != nil {
		log.Fatalf("error while ejecting credentials: %s", err)
	}

	weatherBytes, err := darksky.MakeRequest()
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
