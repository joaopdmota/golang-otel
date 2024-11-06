package usecases_test

import (
	"cep_weather/application/usecases"
	"cep_weather/infra/dtos"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockWeatherRepository struct {
	mock.Mock
}

func (m *MockWeatherRepository) GetWeather(city string) (*dtos.WeatherResponse, error) {
	args := m.Called(city)
	return args.Get(0).(*dtos.WeatherResponse), args.Error(1)
}

func TestSearchByCity_Success(t *testing.T) {
	mockWeatherRepo := new(MockWeatherRepository)
	useCase := usecases.NewWeatherUseCase(mockWeatherRepo)

	exampleWeatherResponse := &dtos.WeatherResponse{
		Location: struct {
			Name           string  `json:"name"`
			Region         string  `json:"region"`
			Country        string  `json:"country"`
			Lat            float64 `json:"lat"`
			Lon            float64 `json:"lon"`
			TzID           string  `json:"tz_id"`
			LocaltimeEpoch int64   `json:"localtime_epoch"`
			Localtime      string  `json:"localtime"`
		}{
			Name:           "New York",
			Region:         "New York",
			Country:        "USA",
			Lat:            40.7128,
			Lon:            -74.0060,
			TzID:           "America/New_York",
			LocaltimeEpoch: 1628290820,
			Localtime:      "2024-11-06 14:00",
		},
		Current: struct {
			LastUpdatedEpoch int64   `json:"last_updated_epoch"`
			LastUpdated      string  `json:"last_updated"`
			TempC            float64 `json:"temp_c"`
			TempF            float64 `json:"temp_f"`
			IsDay            int     `json:"is_day"`
			Condition        struct {
				Text string `json:"text"`
				Icon string `json:"icon"`
				Code int    `json:"code"`
			} `json:"condition"`
			WindMph    float64 `json:"wind_mph"`
			WindKph    float64 `json:"wind_kph"`
			WindDegree int     `json:"wind_degree"`
			WindDir    string  `json:"wind_dir"`
			PressureMb float64 `json:"pressure_mb"`
			PressureIn float64 `json:"pressure_in"`
			PrecipMm   float64 `json:"precip_mm"`
			PrecipIn   float64 `json:"precip_in"`
			Humidity   int     `json:"humidity"`
			Cloud      int     `json:"cloud"`
			FeelslikeC float64 `json:"feelslike_c"`
			FeelslikeF float64 `json:"feelslike_f"`
			WindchillC float64 `json:"windchill_c"`
			WindchillF float64 `json:"windchill_f"`
			HeatindexC float64 `json:"heatindex_c"`
			HeatindexF float64 `json:"heatindex_f"`
			DewpointC  float64 `json:"dewpoint_c"`
			DewpointF  float64 `json:"dewpoint_f"`
			VisKm      float64 `json:"vis_km"`
			VisMiles   float64 `json:"vis_miles"`
			Uv         float64 `json:"uv"`
			GustMph    float64 `json:"gust_mph"`
			GustKph    float64 `json:"gust_kph"`
		}{
			LastUpdatedEpoch: 1628290820,
			LastUpdated:      "2024-11-06 14:00",
			TempC:            22.5,
			TempF:            72.5,
			IsDay:            1,
			Condition: struct {
				Text string `json:"text"`
				Icon string `json:"icon"`
				Code int    `json:"code"`
			}{
				Text: "Clear",
				Icon: "//cdn.weatherapi.com/weather/64x64/day/113.png",
				Code: 1000,
			},
			WindMph:    10.0,
			WindKph:    16.1,
			WindDegree: 180,
			WindDir:    "S",
			PressureMb: 1015,
			PressureIn: 30.03,
			PrecipMm:   0.0,
			PrecipIn:   0.0,
			Humidity:   60,
			Cloud:      0,
			FeelslikeC: 22.0,
			FeelslikeF: 71.6,
			WindchillC: 21.5,
			WindchillF: 70.7,
			HeatindexC: 23.0,
			HeatindexF: 73.4,
			DewpointC:  16.5,
			DewpointF:  61.7,
			VisKm:      10.0,
			VisMiles:   6.2,
			Uv:         5.0,
			GustMph:    15.0,
			GustKph:    24.1,
		},
	}

	mockWeatherRepo.On("GetWeather", "New York").Return(exampleWeatherResponse, nil)

	result, err := useCase.SearchByCity("New York")

	assert.Nil(t, err)
	assert.Equal(t, "22.50", result.TempC)
	assert.Equal(t, "72.50", result.TempF)
	assert.Equal(t, "295.50", result.TempK)
}
