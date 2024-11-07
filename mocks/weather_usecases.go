package mocks

import (
	"cep_weather/application/app"
	"cep_weather/application/usecases/dtos"

	"github.com/stretchr/testify/mock"
)

type MockWeatherUseCase struct {
	mock.Mock
}

func (m *MockWeatherUseCase) SearchByCity(city string) (dtos.WeatherUseCaseOutput, app.Errors) {
	args := m.Called(city)
	return args.Get(0).(dtos.WeatherUseCaseOutput), args.Get(1).(app.Errors)
}

func (m *MockWeatherUseCase) TransformCelsiusToFahrenheit(celsius float64) string {
	args := m.Called(celsius)
	return args.String(0)
}

func (m *MockWeatherUseCase) TransformCelsiusToKelvin(celsius float64) string {
	args := m.Called(celsius)
	return args.String(0)
}
