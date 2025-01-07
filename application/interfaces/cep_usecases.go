package interfaces

import (
	"cep_weather_otel/application/app"
	"cep_weather_otel/infra/dtos"
)

type ICepUseCase interface {
	Search(cep string) (dtos.ViaCepResponse, app.Errors)
	IsEmpty(response dtos.ViaCepResponse) bool
}
