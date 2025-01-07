package repositories_test

import (
	"bytes"
	"cep_weather_otel/infra/dtos"
	"cep_weather_otel/infra/repositories"
	"cep_weather_otel/mocks"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCepRepositoryGetCepSuccess(t *testing.T) {
	viaCepMockResponse := &dtos.ViaCepResponse{
		Cep:         "01001-000",
		Logradouro:  "Praça da Sé",
		Complemento: "lado ímpar",
		Bairro:      "Sé",
		Localidade:  "São Paulo",
		Uf:          "SP",
		Ibge:        "3550308",
	}
	responseBody, _ := json.Marshal(viaCepMockResponse)
	mockResponse := httptest.NewRecorder()
	mockResponse.WriteHeader(http.StatusOK)
	mockResponse.Write(responseBody)

	mockHTTPClient := &mocks.MockHTTPClient{
		Response: mockResponse.Result(),
	}

	repo := repositories.NewCepRepository(mockHTTPClient)

	cep := "01001000"
	result, err := repo.GetCep(cep)

	assert.Nil(t, err)
	assert.Equal(t, viaCepMockResponse, result)
}

func TestCepRepositoryGetCepErrorDecode(t *testing.T) {
	mockResponse := httptest.NewRecorder()
	mockResponse.WriteHeader(http.StatusOK)
	mockResponse.Write(bytes.NewBufferString(`{"}`).Bytes())

	mockHTTPClient := &mocks.MockHTTPClient{
		Response: mockResponse.Result(),
	}

	repo := repositories.NewCepRepository(mockHTTPClient)

	cep := "01001000"
	result, err := repo.GetCep(cep)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestCepRepositoryGetCepError(t *testing.T) {
	mockResponse := httptest.NewRecorder()
	mockResponse.WriteHeader(http.StatusInternalServerError)
	mockResponse.Body = bytes.NewBufferString(`{"error": "internal server error"}`)

	mockClient := &mocks.MockHTTPClient{
		Response: mockResponse.Result(),
	}

	repo := repositories.NewCepRepository(mockClient)

	_, err := repo.GetCep("01001000")

	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "")
	}
}

func TestCepRepositoryGetCepErrorDuringRequest(t *testing.T) {
	mockClient := &mocks.MockHTTPClient{
		Response: nil,
		Err:      errors.New("Failed to fetch"),
	}

	repo := repositories.NewCepRepository(mockClient)

	_, err := repo.GetCep("São Paulo")

	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Failed to fetch")
	}
}
