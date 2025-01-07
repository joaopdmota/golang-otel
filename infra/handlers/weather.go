package handlers

import (
	"cep_weather_otel/application/app"
	"cep_weather_otel/application/interfaces"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

type WeatherHandler struct {
	weatherUseCase interfaces.IWeatherUseCase
}

func NewWeatherHandler(weatherUseCase interfaces.IWeatherUseCase) *WeatherHandler {
	return &WeatherHandler{
		weatherUseCase: weatherUseCase,
	}
}

func (h *WeatherHandler) GetWeather(c echo.Context) error {
	ctx := c.Request().Context()
	tracer := otel.Tracer("city-weather")
	_, span := tracer.Start(ctx, "GetCityWeather", trace.WithAttributes(
		semconv.HTTPMethodKey.String(http.MethodGet),
		semconv.HTTPTargetKey.String(c.Request().URL.Path),
	))
	defer span.End()

	city := c.Param("city")
	span.SetAttributes(semconv.NetPeerNameKey.String(city))
	span.AddEvent("Searching City information")
	cepResponse, err := h.weatherUseCase.SearchByCity(city)

	if err != nil {
		span.AddEvent("Error while fetching city information", trace.WithAttributes(
			semconv.HTTPStatusCodeKey.Int(err[0].Code),
		))
		return c.JSON(err[0].Code, app.CreateErrors(app.Error{
			Code:    err[0].Code,
			Message: err[0].Message,
			Type:    err[0].Type,
		}))
	}

	span.SetAttributes(
		semconv.HTTPStatusCodeKey.Int(http.StatusOK),
		semconv.NetHostNameKey.String(cepResponse.TempC),
	)

	span.AddEvent("Returning final response")

	return c.JSON(http.StatusOK, cepResponse)
}
