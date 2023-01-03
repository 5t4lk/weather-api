package main

import (
	"awesomeProject4/internal/api/darksky"
	"fmt"
	"log"
)

func main() {
	weatherBytes, err := darksky.MakeRequest()
	if err != nil {
		log.Fatalf("error while receiving data from API: %s", err)
	}

	clearWeather, err := darksky.ClearJSON(weatherBytes)
	if err != nil {
		log.Fatalf("error while clearing JSON: %s", err)
	}

	for _, item := range clearWeather.Currently.Icon {
		fmt.Print(string(item))
	}
}
