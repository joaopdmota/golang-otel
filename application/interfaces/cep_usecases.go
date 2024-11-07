package interfaces

import (
	"cep_weather/application/app"
	"cep_weather/infra/dtos"
)

type ICepUseCase interface {
	Search(cep string) (dtos.ViaCepResponse, app.Errors)
	IsEmpty(response dtos.ViaCepResponse) bool
}
