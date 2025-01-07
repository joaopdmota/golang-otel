package usecases

import (
	"cep_weather_otel/application/app"
	"cep_weather_otel/application/interfaces"
	"cep_weather_otel/infra/dtos"
	"net/http"
)

type CepUseCase struct {
	cepRepository     interfaces.ICepRepository
	weatherRepository interfaces.IWeatherRepository
}

func NewCepUseCase(cr interfaces.ICepRepository, wr interfaces.IWeatherRepository) *CepUseCase {
	return &CepUseCase{cepRepository: cr, weatherRepository: wr}
}

func (c *CepUseCase) Search(cep string) (dtos.ViaCepResponse, app.Errors) {
	if len(cep) != 8 {
		return dtos.ViaCepResponse{}, app.CreateErrors(app.Error{
			Message: "Invalid zip code",
			Code:    http.StatusUnprocessableEntity,
			Type:    app.ERROR_UNPROCESSABLE_ENTITY,
		})
	}

	data, err := c.cepRepository.GetCep(cep)

	if err != nil {
		return dtos.ViaCepResponse{}, app.CreateErrors(app.Error{
			Code: http.StatusInternalServerError,
			Type: app.ERROR_UNKNOW,
		})
	}

	if c.IsEmpty(*data) {
		return dtos.ViaCepResponse{}, app.CreateErrors(app.Error{
			Code:    http.StatusNotFound,
			Type:    app.ERROR_NOT_FOUND,
			Message: "Zip code not found",
		})
	}

	return *data, nil
}

func (c *CepUseCase) IsEmpty(response dtos.ViaCepResponse) bool {
	return response.Cep == "" && response.Logradouro == "" && response.Bairro == "" &&
		response.Localidade == "" && response.Uf == ""
}
