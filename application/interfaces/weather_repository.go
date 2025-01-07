package interfaces

import "cep_weather_otel/infra/dtos"

type IWeatherRepository interface {
	GetWeatherApi(city string) (*dtos.WeatherResponse, error)
	GetWeatherMs(cep string) (*dtos.WeatherResponse, error)
}
