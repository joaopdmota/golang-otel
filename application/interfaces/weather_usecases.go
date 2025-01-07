package interfaces

import (
	"cep_weather_otel/application/app"
	"cep_weather_otel/application/usecases/dtos"
)

type IWeatherUseCase interface {
	SearchByCity(city string) (dtos.WeatherUseCaseOutput, app.Errors)
	TransformCelsiusToFahrenheit(celsius float64) string
	TransformCelsiusToKelvin(celsius float64) string
}
