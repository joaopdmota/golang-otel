package interfaces

import (
	"cep_weather/infra/dtos"
)

type ICepRepository interface {
	GetCep(cep string) (*dtos.ViaCepResponse, error)
}
