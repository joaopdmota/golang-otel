package handlers

import (
	"cep_weather/application/app"
	"cep_weather/application/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CepHandler struct {
	cepUseCase     *usecases.CepUseCase
	weatherUseCase *usecases.WeatherUseCase
}

func NewCepHandler(cepUseCase *usecases.CepUseCase, weatherUseCase *usecases.WeatherUseCase) *CepHandler {
	return &CepHandler{
		cepUseCase:     cepUseCase,
		weatherUseCase: weatherUseCase,
	}
}

func (h *CepHandler) GetCEPWeather(c echo.Context) error {
	cep := c.Param("id")

	if cep == "" {
		return c.JSON(http.StatusBadRequest, app.CreateErrors(app.Error{
			Code:    http.StatusBadRequest,
			Type:    app.ERROR_BAD_REQUEST,
			Message: "ID é obrigatório",
		}))
	}

	cepResponse, err := h.cepUseCase.Search(cep)

	if err != nil {
		return c.JSON(http.StatusBadRequest, app.CreateErrors(app.Error{
			Code:    err[0].Code,
			Message: err[0].Message,
			Type:    err[0].Type,
		}))
	}

	weatherResponse, err := h.weatherUseCase.SearchByCity(cepResponse.Localidade)

	if err != nil {
		return c.JSON(http.StatusBadRequest, app.CreateErrors(app.Error{
			Code:    err[0].Code,
			Message: err[0].Message,
			Type:    err[0].Type,
		}))
	}

	return c.JSON(http.StatusOK, weatherResponse)
}
