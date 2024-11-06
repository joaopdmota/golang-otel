package repositories

import (
	"cep_weather/infra/dtos"
	"encoding/json"
	"fmt"
	"net/http"
)

type CepRepository struct{}

func NewCepRepository() *CepRepository {
	return &CepRepository{}
}

func (s *CepRepository) GetCep(cep string) (*dtos.ViaCepResponse, error) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var viaCepResponse dtos.ViaCepResponse
	if err := json.NewDecoder(resp.Body).Decode(&viaCepResponse); err != nil {
		return nil, err
	}

	return &viaCepResponse, nil
}
