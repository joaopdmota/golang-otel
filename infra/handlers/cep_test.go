package handlers

import (
	"cep_weather/application/app"
	"cep_weather/application/usecases/dtos"
	"cep_weather/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	repoDTOS "cep_weather/infra/dtos"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetCEPWeatherHandlerSuccess(t *testing.T) {
	e := echo.New()

	mockCepUseCase := new(mocks.MockCepUseCase)
	mockWeatherUseCase := new(mocks.MockWeatherUseCase)

	cepResponse := repoDTOS.ViaCepResponse{
		Localidade:  "São Paulo",
		Cep:         "01001000",
		Logradouro:  "Praça da Sé",
		Complemento: "lado ímpar",
		Bairro:      "Sé",
		Uf:          "SP",
		Ibge:        "3550308",
		Gia:         "1004",
		Ddd:         "11",
		Siafi:       "7107",
	}

	weatherResponse := dtos.WeatherUseCaseOutput{
		TempC: "25.00",
		TempF: "77.00",
		TempK: "298.15",
	}

	mockCepUseCase.On("Search", "01001000").Return(cepResponse, app.Errors(nil))
	mockWeatherUseCase.On("SearchByCity", "São Paulo").Return(weatherResponse, app.Errors(nil))

	handler := NewCepHandler(mockCepUseCase, mockWeatherUseCase)
	req := httptest.NewRequest(http.MethodGet, "/cep/01001000", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("01001000")
	err := handler.GetCEPWeather(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "25.00")
	assert.Contains(t, rec.Body.String(), "77.00")
	assert.Contains(t, rec.Body.String(), "298.15")
}

func TestGetCEPWeatherHandlerErrorOnCepUseCase(t *testing.T) {
	e := echo.New()

	mockCepUseCase := new(mocks.MockCepUseCase)
	mockWeatherUseCase := new(mocks.MockWeatherUseCase)

	cepResponse := repoDTOS.ViaCepResponse{
		Localidade:  "São Paulo",
		Cep:         "01001000",
		Logradouro:  "Praça da Sé",
		Complemento: "lado ímpar",
		Bairro:      "Sé",
		Uf:          "SP",
		Ibge:        "3550308",
		Gia:         "1004",
		Ddd:         "11",
		Siafi:       "7107",
	}

	errors := app.CreateErrors(app.Error{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
		Type:    app.ERROR_UNKNOW,
	})

	mockCepUseCase.On("Search", "01001000").Return(cepResponse, errors)

	handler := NewCepHandler(mockCepUseCase, mockWeatherUseCase)
	req := httptest.NewRequest(http.MethodGet, "/cep/01001000", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("01001000")

	handler.GetCEPWeather(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestGetCEPWeatherHandlerErrorWeatherUseCase(t *testing.T) {
	e := echo.New()

	mockCepUseCase := new(mocks.MockCepUseCase)
	mockWeatherUseCase := new(mocks.MockWeatherUseCase)

	cepResponse := repoDTOS.ViaCepResponse{
		Localidade:  "São Paulo",
		Cep:         "01001000",
		Logradouro:  "Praça da Sé",
		Complemento: "lado ímpar",
		Bairro:      "Sé",
		Uf:          "SP",
		Ibge:        "3550308",
		Gia:         "1004",
		Ddd:         "11",
		Siafi:       "7107",
	}

	errors := app.CreateErrors(app.Error{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
		Type:    app.ERROR_UNKNOW,
	})

	mockCepUseCase.On("Search", "01001000").Return(cepResponse, app.Errors(nil))
	mockWeatherUseCase.On("SearchByCity", "São Paulo").Return(dtos.WeatherUseCaseOutput{}, errors)

	handler := NewCepHandler(mockCepUseCase, mockWeatherUseCase)
	req := httptest.NewRequest(http.MethodGet, "/cep/01001000", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("01001000")

	handler.GetCEPWeather(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
