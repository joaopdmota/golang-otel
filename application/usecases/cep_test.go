package usecases_test

import (
	"cep_weather/application/app"
	"cep_weather/application/usecases"
	"cep_weather/infra/dtos"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCepRepository struct {
	mock.Mock
}

func (m *MockCepRepository) GetCep(cep string) (*dtos.ViaCepResponse, error) {
	args := m.Called(cep)
	return args.Get(0).(*dtos.ViaCepResponse), args.Error(1)
}

func TestSearch_InvalidCep(t *testing.T) {
	mockCepRepo := new(MockCepRepository)
	useCase := usecases.NewCepUseCase(mockCepRepo)
	cep := "12345"

	mockCepRepo.On("GetCep", cep).Return(dtos.ViaCepResponse{}, app.CreateErrors(
		app.Error{
			Code: http.StatusUnprocessableEntity,
			Type: app.ERROR_UNPROCESSABLE_ENTITY,
		},
	))

	result, err := useCase.Search(cep)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, err[0].Code)
	assert.Equal(t, app.ERROR_UNPROCESSABLE_ENTITY, err[0].Type)
	assert.Empty(t, result)
}

func TestSearch_CepError(t *testing.T) {
	mockCepRepo := new(MockCepRepository)
	useCase := usecases.NewCepUseCase(mockCepRepo)

	cep := "12345678"

	mockCepRepo.On("GetCep", cep).Return(&dtos.ViaCepResponse{}, errors.New("Unknown error"))

	result, err := useCase.Search(cep)

	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, err[0].Code)
	assert.Equal(t, app.ERROR_UNKNOW, err[0].Type)
	assert.Empty(t, result)
}

func TestSearch_CepSuccess(t *testing.T) {
	mockCepRepo := new(MockCepRepository)
	useCase := usecases.NewCepUseCase(mockCepRepo)

	cep := "12345678"

	mockCepRepo.On("GetCep", cep).Return(&dtos.ViaCepResponse{
		Cep:         "11702-150",
		Logradouro:  "Rua Chile",
		Complemento: "",
		Bairro:      "Guilhermina",
		Localidade:  "Praia Grande",
		Uf:          "SP",
		Ibge:        "3541000",
		Gia:         "5587",
		Ddd:         "13",
		Siafi:       "6921",
	}, nil)

	result, err := useCase.Search(cep)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestSearch_CepNotFound(t *testing.T) {
	mockCepRepo := new(MockCepRepository)
	useCase := usecases.NewCepUseCase(mockCepRepo)

	cep := "12345678"

	mockCepRepo.On("GetCep", cep).Return(&dtos.ViaCepResponse{}, nil)

	_, err := useCase.Search(cep)

	assert.Equal(t, http.StatusNotFound, err[0].Code)
}

func TestIsEmpty(t *testing.T) {
	useCase := usecases.NewCepUseCase(nil)

	emptyResponse := dtos.ViaCepResponse{}

	assert.True(t, useCase.IsEmpty(emptyResponse))
}
