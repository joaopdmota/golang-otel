package usecases

import (
	"cep_weather/application/app"
	"cep_weather/application/interfaces"
	"cep_weather/infra/dtos"
	"net/http"
)

type CepUseCase struct {
	cepRepository interfaces.ICepRepository
}

func NewCepUseCase(cr interfaces.ICepRepository) *CepUseCase {
	return &CepUseCase{cepRepository: cr}
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

	if isEmpty(*data) {
		return dtos.ViaCepResponse{}, app.CreateErrors(app.Error{
			Code:    http.StatusNotFound,
			Type:    app.ERROR_NOT_FOUND,
			Message: "Zip code not found",
		})
	}

	return *data, nil
}

func isEmpty(response dtos.ViaCepResponse) bool {
	return response.Cep == "" && response.Logradouro == "" && response.Bairro == "" &&
		response.Localidade == "" && response.Uf == ""
}
