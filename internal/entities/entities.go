package entities

type General struct {
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timezone  string    `json:"timezone"`
	Currently Currently `json:"currently"`
	Minutely  Minutely  `json:"minutely"`
	Hourly    Hourly    `json:"hourly"`
	Daily     Daily     `json:"daily"`
	Flags     Flags     `json:"flags"`
	Offset    int64     `json:"offset"`
}

type Currently struct {
	Time                 int64    `json:"time"`
	Summary              string   `json:"summary"`
	Icon                 Icon     `json:"icon"`
	NearestStormDistance *int64   `json:"nearestStormDistance,omitempty"`
	PrecipIntensity      float64  `json:"precipIntensity"`
	PrecipIntensityError *float64 `json:"precipIntensityError,omitempty"`
	PrecipProbability    float64  `json:"precipProbability"`
	PrecipType           *Icon    `json:"precipType,omitempty"`
	Temperature          float64  `json:"temperature"`
	ApparentTemperature  float64  `json:"apparentTemperature"`
	DewPoint             float64  `json:"dewPoint"`
	Humidity             float64  `json:"humidity"`
	Pressure             float64  `json:"pressure"`
	WindSpeed            float64  `json:"windSpeed"`
	WindGust             float64  `json:"windGust"`
	WindBearing          int64    `json:"windBearing"`
	CloudCover           float64  `json:"cloudCover"`
	UvIndex              int64    `json:"uvIndex"`
	Visibility           float64  `json:"visibility"`
	Ozone                float64  `json:"ozone"`
}

type Daily struct {
	Summary string       `json:"summary"`
	Icon    Icon         `json:"icon"`
	Data    []DailyDatum `json:"data"`
}

type DailyDatum struct {
	Time                        int64   `json:"time"`
	Summary                     string  `json:"summary"`
	Icon                        string  `json:"icon"`
	SunriseTime                 int64   `json:"sunriseTime"`
	SunsetTime                  int64   `json:"sunsetTime"`
	MoonPhase                   float64 `json:"moonPhase"`
	PrecipIntensity             float64 `json:"precipIntensity"`
	PrecipIntensityMax          float64 `json:"precipIntensityMax"`
	PrecipIntensityMaxTime      int64   `json:"precipIntensityMaxTime"`
	PrecipProbability           float64 `json:"precipProbability"`
	PrecipType                  Icon    `json:"precipType"`
	TemperatureHigh             float64 `json:"temperatureHigh"`
	TemperatureHighTime         int64   `json:"temperatureHighTime"`
	TemperatureLow              float64 `json:"temperatureLow"`
	TemperatureLowTime          int64   `json:"temperatureLowTime"`
	ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh"`
	ApparentTemperatureHighTime int64   `json:"apparentTemperatureHighTime"`
	ApparentTemperatureLow      float64 `json:"apparentTemperatureLow"`
	ApparentTemperatureLowTime  int64   `json:"apparentTemperatureLowTime"`
	DewPoint                    float64 `json:"dewPoint"`
	Humidity                    float64 `json:"humidity"`
	Pressure                    float64 `json:"pressure"`
	WindSpeed                   float64 `json:"windSpeed"`
	WindGust                    float64 `json:"windGust"`
	WindGustTime                int64   `json:"windGustTime"`
	WindBearing                 int64   `json:"windBearing"`
	CloudCover                  float64 `json:"cloudCover"`
	UvIndex                     int64   `json:"uvIndex"`
	UvIndexTime                 int64   `json:"uvIndexTime"`
	Visibility                  float64 `json:"visibility"`
	Ozone                       float64 `json:"ozone"`
	TemperatureMin              float64 `json:"temperatureMin"`
	TemperatureMinTime          int64   `json:"temperatureMinTime"`
	TemperatureMax              float64 `json:"temperatureMax"`
	TemperatureMaxTime          int64   `json:"temperatureMaxTime"`
	ApparentTemperatureMin      float64 `json:"apparentTemperatureMin"`
	ApparentTemperatureMinTime  int64   `json:"apparentTemperatureMinTime"`
	ApparentTemperatureMax      float64 `json:"apparentTemperatureMax"`
	ApparentTemperatureMaxTime  int64   `json:"apparentTemperatureMaxTime"`
}

type Flags struct {
	Sources           []string `json:"sources"`
	MeteoalarmLicense string   `json:"meteoalarm-license"`
	NearestStation    float64  `json:"nearest-station"`
	Units             string   `json:"units"`
}

type Hourly struct {
	Summary string      `json:"summary"`
	Icon    Icon        `json:"icon"`
	Data    []Currently `json:"data"`
}

type Minutely struct {
	Summary string          `json:"summary"`
	Icon    Icon            `json:"icon"`
	Data    []MinutelyDatum `json:"data"`
}

type MinutelyDatum struct {
	Time                 int64   `json:"time"`
	PrecipIntensity      float64 `json:"precipIntensity"`
	PrecipIntensityError float64 `json:"precipIntensityError"`
	PrecipProbability    float64 `json:"precipProbability"`
	PrecipType           Icon    `json:"precipType"`
}

type Icon string

const (
	ClearNight        Icon = "clear-night"
	Cloudy            Icon = "cloudy"
	PartlyCloudyNight Icon = "partly-cloudy-night"
	Rain              Icon = "rain"
)
