package interfaces

import "cep_weather/infra/dtos"

type IWeatherRepository interface {
	GetWeather(city string) (*dtos.WeatherResponse, error)
}
