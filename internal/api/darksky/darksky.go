package darksky

import (
	"awesomeProject4/internal/entities"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	url           = "https://dark-sky.p.rapidapi.com/50.9574,0.5107?units=auto"
	xRapidAPIKey  = "85e69f9465msh9b8963f3d02fc11p1fbbc9jsn64207893b164"
	xRapidAPIHost = "dark-sky.p.rapidapi.com"
)

func MakeRequest() ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-RapidAPI-Key", xRapidAPIKey)
	req.Header.Add("X-RapidAPI-Host", xRapidAPIHost)

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
