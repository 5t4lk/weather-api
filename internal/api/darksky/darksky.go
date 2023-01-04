package darksky

import (
	"awesomeProject4/internal/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func MakeRequest(latitude, longitude, apiKey, apiHost string) ([]byte, error) {
	url := fmt.Sprintf("https://" + apiHost + "/" + latitude + "," + longitude + "?units=auto&lang=en")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", apiHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ClearJSON(messyData []byte) (*entities.General, error) {
	var data entities.General

	err := json.Unmarshal(messyData, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
