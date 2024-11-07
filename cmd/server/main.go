package main

import (
	"cep_weather/application/usecases"
	config "cep_weather/infra/config"
	"cep_weather/infra/handlers"
	repositories "cep_weather/infra/repositories"
	httpclient "cep_weather/infra/repositories/http"
	"cep_weather/infra/server"
	"fmt"
	"net/http"
	"time"
)

func main() {
	config.LoadEnvs()

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	httpClient := httpclient.NewDefaultHTTPClient(client)

	service := server.NewHTTPService("8080")
	cepRepository := repositories.NewCepRepository(httpClient)
	weatherRepository := repositories.NewWeatherRepository(httpClient)

	cepUseCase := usecases.NewCepUseCase(cepRepository)
	weatherUseCase := usecases.NewWeatherUseCase(weatherRepository)

	cepHandler := handlers.NewCepHandler(cepUseCase, weatherUseCase)

	service.AddRoute(http.MethodGet, "/cep/:id", cepHandler.GetCEPWeather)

	if err := service.Start(); err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}

	fmt.Println("Running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
