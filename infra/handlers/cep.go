package handlers

import (
	"cep_weather_otel/application/app"
	"cep_weather_otel/application/interfaces"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

type CepHandler struct {
	cepUseCase        interfaces.ICepUseCase
	weatherUseCase    interfaces.IWeatherUseCase
	weatherRepository interfaces.IWeatherRepository
}

type GetCepWeatherOutput struct {
	City  string `json:"city"`
	TempC string `json:"temp_C"`
	TempF string `json:"temp_F"`
	TempK string `json:"temp_K"`
}

func NewCepHandler(cepUseCase interfaces.ICepUseCase, weatherUseCase interfaces.IWeatherUseCase, weatherRepository interfaces.IWeatherRepository) *CepHandler {
	return &CepHandler{
		cepUseCase:        cepUseCase,
		weatherRepository: weatherRepository,
		weatherUseCase:    weatherUseCase,
	}
}

func (h *CepHandler) GetCEPWeather(c echo.Context) error {
	ctx := c.Request().Context()
	tracer := otel.Tracer("cep-weather")
	_, span := tracer.Start(ctx, "GetCEPWeather", trace.WithAttributes(
		semconv.HTTPMethodKey.String(http.MethodGet),
		semconv.HTTPTargetKey.String(c.Request().URL.Path),
	))
	defer span.End()

	cep := c.Param("id")
	span.SetAttributes(semconv.NetPeerNameKey.String(cep))

	span.AddEvent("Searching CEP information")
	cepResponse, cepErr := h.cepUseCase.Search(cep)

	if cepErr != nil {
		span.AddEvent("Error while fetching CEP information", trace.WithAttributes(
			semconv.HTTPStatusCodeKey.Int(cepErr[0].Code),
		))
		return c.JSON(cepErr[0].Code, app.CreateErrors(app.Error{
			Code:    cepErr[0].Code,
			Message: cepErr[0].Message,
			Type:    cepErr[0].Type,
		}))
	}

	span.AddEvent("Fetching weather information")
	weatherResponse, weatherErr := h.weatherRepository.GetWeatherMs(cepResponse.Localidade)

	if weatherErr != nil {
		span.AddEvent("Error while fetching weather information", trace.WithAttributes(
			semconv.HTTPStatusCodeKey.Int(http.StatusInternalServerError),
		))
		return c.JSON(echo.ErrInternalServerError.Code, app.CreateErrors(app.Error{
			Code: http.StatusInternalServerError,
			Type: app.ERROR_UNKNOW,
		}))
	}

	span.AddEvent("Transforming weather data")
	response := GetCepWeatherOutput{
		City:  cepResponse.Localidade,
		TempC: fmt.Sprintf("%.2f", weatherResponse.Current.TempC),
		TempF: h.weatherUseCase.TransformCelsiusToFahrenheit(weatherResponse.Current.TempC),
		TempK: h.weatherUseCase.TransformCelsiusToKelvin(weatherResponse.Current.TempC),
	}

	span.SetAttributes(
		semconv.HTTPStatusCodeKey.Int(http.StatusOK),
		semconv.NetHostNameKey.String(response.City),
	)

	// Retorna a resposta
	span.AddEvent("Returning final response")
	return c.JSON(http.StatusOK, response)
}
