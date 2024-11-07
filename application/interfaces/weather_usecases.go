package interfaces

import (
	"cep_weather/application/app"
	"cep_weather/application/usecases/dtos"
)

type IWeatherUseCase interface {
	SearchByCity(city string) (dtos.WeatherUseCaseOutput, app.Errors)
	TransformCelsiusToFahrenheit(celsius float64) string
	TransformCelsiusToKelvin(celsius float64) string
}
