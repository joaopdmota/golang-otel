package main

import (
	"cep_weather_otel/application/usecases"
	config "cep_weather_otel/infra/config"
	"cep_weather_otel/infra/handlers"
	repositories "cep_weather_otel/infra/repositories"
	httpclient "cep_weather_otel/infra/repositories/http"
	"cep_weather_otel/infra/server"
	"cep_weather_otel/pkg"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

func startServer(port string, routes func(*server.HTTPService)) error {
	service := server.NewHTTPService(port)

	routes(service)

	fmt.Printf("Running on http://localhost:%s\n", port)
	return service.Start()
}

func main() {
	config.LoadEnvs()

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	httpClient := httpclient.NewDefaultHTTPClient(client)

	cepRepository := repositories.NewCepRepository(httpClient)
	weatherRepository := repositories.NewWeatherRepository(httpClient)

	cepUseCase := usecases.NewCepUseCase(cepRepository, weatherRepository)
	weatherUseCase := usecases.NewWeatherUseCase(weatherRepository)

	cepHandler := handlers.NewCepHandler(cepUseCase, weatherUseCase, weatherRepository)
	weatherHandler := handlers.NewWeatherHandler(weatherUseCase)

	errChan := make(chan error, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		errChan <- startServer("8080", func(service *server.HTTPService) {
			service.AddRoute(http.MethodGet, "/cep/:id", cepHandler.GetCEPWeather)
		})
	}()

	go func() {
		defer wg.Done()
		errChan <- startServer("8081", func(service *server.HTTPService) {
			service.AddRoute(http.MethodGet, "/weather/:city", weatherHandler.GetWeather)
		})
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	_, err := pkg.SetupOTelSDK(ctx)

	if err != nil {
		fmt.Println("Erro ao iniciar observabilidade:", err)
	}

	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	for err := range errChan {
		if err != nil {
			fmt.Println("Erro ao iniciar um dos servidores:", err)
		}
	}
}
