package handlers

import (
	"cep_weather/application/app"
	"cep_weather/application/interfaces"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CepHandler struct {
	cepUseCase     interfaces.ICepUseCase
	weatherUseCase interfaces.IWeatherUseCase
}

func NewCepHandler(cepUseCase interfaces.ICepUseCase, weatherUseCase interfaces.IWeatherUseCase) *CepHandler {
	return &CepHandler{
		cepUseCase:     cepUseCase,
		weatherUseCase: weatherUseCase,
	}
}

func (h *CepHandler) GetCEPWeather(c echo.Context) error {
	cep := c.Param("id")

	cepResponse, err := h.cepUseCase.Search(cep)

	if err != nil {
		return c.JSON(err[0].Code, app.CreateErrors(app.Error{
			Code:    err[0].Code,
			Message: err[0].Message,
			Type:    err[0].Type,
		}))
	}

	weatherResponse, err := h.weatherUseCase.SearchByCity(cepResponse.Localidade)

	if err != nil {
		return c.JSON(err[0].Code, app.CreateErrors(app.Error{
			Code:    err[0].Code,
			Message: err[0].Message,
			Type:    err[0].Type,
		}))
	}

	return c.JSON(http.StatusOK, weatherResponse)
}
