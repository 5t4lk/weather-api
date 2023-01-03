package user

import (
	"awesomeProject4/internal/entities"
	"fmt"
)

func Output(data *entities.General) error {
	fmt.Printf("Your latitude: %.4f\n", data.Latitude)
	fmt.Printf("Your longitude: %.4f\n", data.Longitude)
	fmt.Printf("Timezone: %s\n", data.Timezone)
	fmt.Print("___\n")
	fmt.Printf("Time: %d\n", data.Currently.Time)
	fmt.Printf("Summary: %s\n", data.Currently.Summary)
	fmt.Printf("Icon: %s\n", data.Currently.Icon)
	fmt.Printf("Precipitation intensity: %.2f\n", data.Currently.PrecipIntensity)
	fmt.Printf("Precipitation probability: %.2f\n", data.Currently.PrecipProbability)
	fmt.Printf("Precipitation type: %s\n", data.Currently.PrecipType)
	fmt.Printf("Temperature: %.2f\n", data.Currently.Temperature)
	fmt.Printf("Apparent temperature: %.2f\n", data.Currently.ApparentTemperature)
	fmt.Printf("Dew point: %.2f\n", data.Currently.DewPoint)
	fmt.Printf("Humidity: %.2f\n", data.Currently.Humidity)
	fmt.Printf("Pressure: %.2f\n", data.Currently.Pressure)
	fmt.Printf("Wind speed: %.2f\n", data.Currently.WindSpeed)
	fmt.Printf("Wind gust: %.2f\n", data.Currently.WindGust)
	fmt.Printf("Wind bearing: %d\n", data.Currently.WindBearing)
	fmt.Printf("Cloud cover: %.2f\n", data.Currently.CloudCover)
	fmt.Printf("UV index: %d\n", data.Currently.UvIndex)
	fmt.Printf("Visibility: %d\n", data.Currently.Visibility)
	fmt.Printf("Ozone: %.2f\n", data.Currently.Ozone)
	return nil
}
