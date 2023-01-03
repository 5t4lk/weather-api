package entities

type General struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Currently Currently
	Hourly    Hourly
	Daily     Daily
	Flags     Flags
	Offset    int `json:"offset"`
}

type Currently struct {
	Time                int     `json:"time"`
	Summary             string  `json:"summary"`
	Icon                string  `json:"icon"`
	PrecipIntensity     float64 `json:"precipIntensity"`
	PrecipProbability   float64 `json:"precipProbability"`
	PrecipType          string  `json:"precipType"`
	Temperature         float64 `json:"temperature"`
	ApparentTemperature float64 `json:"apparentTemperature"`
	DewPoint            float64 `json:"dewPoint"`
	Humidity            float64 `json:"humidity"`
	Pressure            float64 `json:"pressure"`
	WindSpeed           float64 `json:"windSpeed"`
	WindGust            float64 `json:"windGust"`
	WindBearing         int     `json:"windBearing"`
	CloudCover          float64 `json:"cloudCover"`
	UvIndex             int     `json:"uvIndex"`
	Visibility          int     `json:"visibility"`
	Ozone               float64 `json:"ozone"`
}

type Hourly struct {
	Summary string `json:"summary"`
	Icon    string `json:"icon"`
	Data    []struct {
		Time                int     `json:"time"`
		Summary             string  `json:"summary"`
		Icon                string  `json:"icon"`
		PrecipIntensity     float64 `json:"precipIntensity"`
		PrecipProbability   float64 `json:"precipProbability"`
		PrecipType          string  `json:"precipType,omitempty"`
		Temperature         float64 `json:"temperature"`
		ApparentTemperature float64 `json:"apparentTemperature"`
		DewPoint            float64 `json:"dewPoint"`
		Humidity            float64 `json:"humidity"`
		Pressure            float64 `json:"pressure"`
		WindSpeed           float64 `json:"windSpeed"`
		WindGust            float64 `json:"windGust"`
		WindBearing         int     `json:"windBearing"`
		CloudCover          float64 `json:"cloudCover"`
		UvIndex             int     `json:"uvIndex"`
		Visibility          float64 `json:"visibility"`
		Ozone               float64 `json:"ozone"`
	} `json:"data"`
}

type Daily struct {
	Summary string `json:"summary"`
	Icon    string `json:"icon"`
	Data    []struct {
		Time                        int     `json:"time"`
		Summary                     string  `json:"summary"`
		Icon                        string  `json:"icon"`
		SunriseTime                 int     `json:"sunriseTime"`
		SunsetTime                  int     `json:"sunsetTime"`
		MoonPhase                   float64 `json:"moonPhase"`
		PrecipIntensity             float64 `json:"precipIntensity"`
		PrecipIntensityMax          float64 `json:"precipIntensityMax"`
		PrecipIntensityMaxTime      int     `json:"precipIntensityMaxTime"`
		PrecipProbability           float64 `json:"precipProbability"`
		PrecipType                  string  `json:"precipType"`
		TemperatureHigh             float64 `json:"temperatureHigh"`
		TemperatureHighTime         int     `json:"temperatureHighTime"`
		TemperatureLow              float64 `json:"temperatureLow"`
		TemperatureLowTime          int     `json:"temperatureLowTime"`
		ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh"`
		ApparentTemperatureHighTime int     `json:"apparentTemperatureHighTime"`
		ApparentTemperatureLow      float64 `json:"apparentTemperatureLow"`
		ApparentTemperatureLowTime  int     `json:"apparentTemperatureLowTime"`
		DewPoint                    float64 `json:"dewPoint"`
		Humidity                    float64 `json:"humidity"`
		Pressure                    float64 `json:"pressure"`
		WindSpeed                   float64 `json:"windSpeed"`
		WindGust                    float64 `json:"windGust"`
		WindGustTime                int     `json:"windGustTime"`
		WindBearing                 int     `json:"windBearing"`
		CloudCover                  float64 `json:"cloudCover"`
		UvIndex                     int     `json:"uvIndex"`
		UvIndexTime                 int     `json:"uvIndexTime"`
		Visibility                  float64 `json:"visibility"`
		Ozone                       float64 `json:"ozone"`
		TemperatureMin              float64 `json:"temperatureMin"`
		TemperatureMinTime          int     `json:"temperatureMinTime"`
		TemperatureMax              float64 `json:"temperatureMax"`
		TemperatureMaxTime          int     `json:"temperatureMaxTime"`
		ApparentTemperatureMin      float64 `json:"apparentTemperatureMin"`
		ApparentTemperatureMinTime  int     `json:"apparentTemperatureMinTime"`
		ApparentTemperatureMax      float64 `json:"apparentTemperatureMax"`
		ApparentTemperatureMaxTime  int     `json:"apparentTemperatureMaxTime"`
	} `json:"data"`
}

type Flags struct {
	Sources            []string `json:"sources"`
	MeteoalarmLicense  string   `json:"meteoalarm-license"`
	NearestStation     float64  `json:"nearest-station"`
	Units              string   `json:"units"`
	DarkskyUnavailable string   `json:"darksky-unavailable"`
}
