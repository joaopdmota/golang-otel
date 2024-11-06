package main

import (
	"cep_weather/application/usecases"
	config "cep_weather/infra/config"
	"cep_weather/infra/handlers"
	repositories "cep_weather/infra/repositories"
	"cep_weather/infra/server"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	_, err := config.LoadEnvs()
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	service := server.NewHTTPService("8080")
	cepRepository := repositories.NewCepRepository()
	weatherRepository := repositories.NewWeatherRepository()

	cepUseCase := usecases.NewCepUseCase(cepRepository)
	weatherUseCase := usecases.NewWeatherUseCase(weatherRepository)

	cepHandler := handlers.NewCepHandler(cepUseCase, weatherUseCase)

	service.AddRoute(http.MethodGet, "/cep/:id", cepHandler.GetCEPWeather)

	if err := service.Start(); err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}

	http.HandleFunc("/", handler)
	fmt.Println("Running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
