package settings

import (
	"bufio"
	"fmt"
	"os"
)

func ScanFile(name string) (array []string, err error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		array = append(array, scanner.Text())
	}

	return array, nil
}

func EjectData(array []string) (string, string, string, string, error) {
	var latitude, longitude, apiKey, apiHost string

	for number, value := range array {
		switch number {
		case 0:
			latitude = value
		case 1:
			longitude = value
		case 2:
			apiKey = value
		case 3:
			apiHost = value
		default:
			return "", "", "", "", fmt.Errorf("wrong ranging file")
		}
	}

	return latitude, longitude, apiKey, apiHost, nil
}
