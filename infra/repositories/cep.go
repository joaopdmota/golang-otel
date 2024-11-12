package repositories

import (
	"cep_weather/infra/dtos"
	http "cep_weather/infra/repositories/http"
	"encoding/json"
	"fmt"
	httpNet "net/http"
)

type CepRepository struct {
	client http.HTTPClient
}

func NewCepRepository(client http.HTTPClient) *CepRepository {
	return &CepRepository{client: client}
}

func (s *CepRepository) GetCep(cep string) (*dtos.ViaCepResponse, error) {
	resp, err := s.client.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= httpNet.StatusBadRequest {
		return nil, fmt.Errorf("erro na requisição: status code %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	var viaCepResponse dtos.ViaCepResponse
	if err := json.NewDecoder(resp.Body).Decode(&viaCepResponse); err != nil {
		return nil, err
	}

	return &viaCepResponse, nil
}
