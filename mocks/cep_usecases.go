package mocks

import (
	"cep_weather_otel/application/app"
	"cep_weather_otel/infra/dtos"

	"github.com/stretchr/testify/mock"
)

type MockCepUseCase struct {
	mock.Mock
}

func (m *MockCepUseCase) Search(cep string) (dtos.ViaCepResponse, app.Errors) {
	args := m.Called(cep)
	return args.Get(0).(dtos.ViaCepResponse), args.Get(1).(app.Errors)
}

func (m *MockCepUseCase) IsEmpty(response dtos.ViaCepResponse) bool {
	args := m.Called(response)
	return args.Bool(0)
}
