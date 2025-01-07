package interfaces

import (
	"cep_weather_otel/infra/dtos"
)

type ICepRepository interface {
	GetCep(cep string) (*dtos.ViaCepResponse, error)
}
