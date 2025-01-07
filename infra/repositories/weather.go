package repositories

import (
	env "cep_weather_otel/infra/config"
	"cep_weather_otel/infra/dtos"
	http "cep_weather_otel/infra/repositories/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
	httpNet "net/http"
	"net/url"
)

type WeatherRepository struct {
	client http.HTTPClient
}

func NewWeatherRepository(client http.HTTPClient) *WeatherRepository {
	return &WeatherRepository{client: client}
}

func (s *WeatherRepository) GetWeatherMs(city string) (*dtos.WeatherResponse, error) {
	encodedCity := url.QueryEscape(city)

	url := fmt.Sprintf("%s/weather/%s", env.Config.WeatherMicroserviceUrl, encodedCity)

	resp, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= httpNet.StatusBadRequest {
		return nil, fmt.Errorf("erro na requisição: status code %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weather dtos.WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, err
	}

	return &weather, nil
}

func (s *WeatherRepository) GetWeatherApi(city string) (*dtos.WeatherResponse, error) {
	encodedCity := url.QueryEscape(city)

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", env.Config.WeatherApiKey, encodedCity)

	resp, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= httpNet.StatusBadRequest {
		return nil, fmt.Errorf("erro na requisição: status code %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weather dtos.WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, err
	}

	return &weather, nil
}
