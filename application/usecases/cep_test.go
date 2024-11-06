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
