package repositories_test

import (
	"bytes"
	env "cep_weather/infra/config"
	"cep_weather/infra/dtos"
	"cep_weather/infra/repositories"
	"cep_weather/mocks"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	httpNet "net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Setenv("WEATHER_API_KEY", "123")
	os.Setenv("API_PORT", "8080")
	env.LoadEnvs()
	code := m.Run()
	os.Exit(code)
}

func TestWeatherRepositoryGetWeatherSuccess(t *testing.T) {
	viaCepMockResponse := &dtos.WeatherResponse{
		Location: dtos.Location{
			Name:   "São Paulo",
			Region: "SP",
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

	responseBody, _ := json.Marshal(viaCepMockResponse)
	mockResponse := httptest.NewRecorder()
	mockResponse.WriteHeader(http.StatusOK)
	mockResponse.Write(responseBody)

	mockClient := &mocks.MockHTTPClient{
		Response: mockResponse.Result(),
	}

	repo := repositories.NewWeatherRepository(mockClient)

	weather, _ := repo.GetWeather("São Paulo")

	assert.NotNil(t, weather)
	assert.Equal(t, "São Paulo", weather.Location.Name)
	assert.Equal(t, "SP", weather.Location.Region)
	assert.Equal(t, 22.5, weather.Current.TempC)
	assert.Equal(t, "Partly cloudy", weather.Current.Condition.Text)
}

func TestWeatherRepositoryGetWeatherError(t *testing.T) {
	mockResponse := httptest.NewRecorder()
	mockResponse.WriteHeader(httpNet.StatusInternalServerError)
	mockResponse.Body = bytes.NewBufferString(`{"error": "internal server error"}`)

	mockClient := &mocks.MockHTTPClient{
		Response: mockResponse.Result(),
	}

	repo := repositories.NewWeatherRepository(mockClient)

	_, err := repo.GetWeather("São Paulo")

	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "")
	}
}

func TestWeatherRepositoryGetWeatherErrorDuringRequest(t *testing.T) {
	mockClient := &mocks.MockHTTPClient{
		Response: nil,
		Err:      errors.New("Failed to fetch"),
	}

	repo := repositories.NewWeatherRepository(mockClient)

	_, err := repo.GetWeather("São Paulo")

	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Failed to fetch")
	}
}

func TestWeatherRepositoryGetWeatherErrorDuringDecodeResponse(t *testing.T) {
	mockResponse := httptest.NewRecorder()
	mockResponse.WriteHeader(http.StatusOK)
	mockResponse.Write(bytes.NewBufferString(``).Bytes())

	mockClient := &mocks.MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(&mocks.MockErrorReader{}),
		},
	}

	repo := repositories.NewWeatherRepository(mockClient)

	_, err := repo.GetWeather("São Paulo")

	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "failed to read response body")
	}
}

func TestWeatherRepositoryGetWeatherErrorDuringDecodeBody(t *testing.T) {
	mockResponse := httptest.NewRecorder()
	mockResponse.WriteHeader(http.StatusOK)
	mockResponse.Write(bytes.NewBufferString(`{"}`).Bytes())

	mockClient := &mocks.MockHTTPClient{
		Response: mockResponse.Result(),
	}

	repo := repositories.NewWeatherRepository(mockClient)

	_, err := repo.GetWeather("São Paulo")

	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "")
	}
}
