package repositories

import (
	env "cep_weather/infra/config"
	"cep_weather/infra/dtos"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type WeatherRepository struct{}

func NewWeatherRepository() *WeatherRepository {
	return &WeatherRepository{}
}

func (s *WeatherRepository) GetWeather(city string) (*dtos.WeatherResponse, error) {
	encodedCity := url.QueryEscape(city)

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", env.Config.APIKey, encodedCity)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
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
