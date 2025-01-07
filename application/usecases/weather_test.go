package usecases_test

import (
	"cep_weather_otel/application/app"
	"cep_weather_otel/application/usecases"
	"cep_weather_otel/infra/dtos"
	"errors"
	"net/http"
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

func TestSearchByCity_Error(t *testing.T) {
	mockWeatherRepo := new(MockWeatherRepository)
	useCase := usecases.NewWeatherUseCase(mockWeatherRepo)

	mockWeatherRepo.On("GetWeather", "New York").Return(&dtos.WeatherResponse{}, errors.New("Unknown error"))

	_, err := useCase.SearchByCity("New York")

	assert.Equal(t, err, app.CreateErrors(app.Error{
		Code: http.StatusUnprocessableEntity,
		Type: app.ERROR_UNKNOW,
	}))
}

func TestSearchByCity_Success(t *testing.T) {
	mockWeatherRepo := new(MockWeatherRepository)
	useCase := usecases.NewWeatherUseCase(mockWeatherRepo)

	exampleWeatherResponse := &dtos.WeatherResponse{
		Location: dtos.Location{
			Name:   "New York",
			Region: "USA",
		},
		Current: dtos.Current{
			TempC:     22.5,
			Condition: dtos.Condition{Text: "Partly cloudy"},
		},
		WindMph:    5.0,
		WindKph:    8.0,
		WindDegree: 180,
		WindDir:    "S",
		PressureMb: 1010.0,
		PrecipMm:   0.0,
		Humidity:   50,
		Cloud:      25,
		FeelslikeC: 22.0,
		VisKm:      10.0,
		Uv:         5.0,
		GustMph:    10.0,
		GustKph:    16.1,
		PressureIn: 30.03,
		PrecipIn:   0.0,
		FeelslikeF: 71.6,
		VisMiles:   6.2,
		HeatindexC: 23.0,
		HeatindexF: 73.4,
		DewpointC:  16.5,
		DewpointF:  61.7,
		WindchillC: 21.5,
		WindchillF: 70.7,
	}

	mockWeatherRepo.On("GetWeather", "New York").Return(exampleWeatherResponse, nil)

	result, err := useCase.SearchByCity("New York")

	assert.Nil(t, err)
	assert.Equal(t, "22.50", result.TempC)
	assert.Equal(t, "72.50", result.TempF)
	assert.Equal(t, "295.50", result.TempK)
}

func TestTransformCelsiusToFahrenheit(t *testing.T) {
	mockWeatherRepo := new(MockWeatherRepository)
	useCase := usecases.NewWeatherUseCase(mockWeatherRepo)
	fahrenheit := useCase.TransformCelsiusToFahrenheit(22.5)

	assert.Equal(t, "72.50", fahrenheit)
}

func TestTransformCelsiusToKelvin(t *testing.T) {
	mockWeatherRepo := new(MockWeatherRepository)
	useCase := usecases.NewWeatherUseCase(mockWeatherRepo)

	kelvin := useCase.TransformCelsiusToKelvin(22.5)

	assert.Equal(t, "295.50", kelvin)
}
