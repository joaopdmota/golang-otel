package env

import (
	"fmt"
	"os"
	"strconv"
)

type ConfigMap struct {
	WeatherApiKey           string
	ApiPort                 int
	CepMicroserviceUrl      string
	CepMicroserviceName     string
	WeatherMicroserviceUrl  string
	WeatherMicroserviceName string
}

var Config *ConfigMap

func LoadEnvs() *ConfigMap {
	if Config != nil {
		return Config
	}

	Config = &ConfigMap{
		WeatherApiKey:           GetEnvString("WEATHER_API_KEY"),
		ApiPort:                 GetEnvNumber("API_PORT"),
		CepMicroserviceUrl:      GetEnvString("CEP_MICROSERVICE_URL"),
		CepMicroserviceName:     GetEnvString("CEP_MICROSERVICE_NAME"),
		WeatherMicroserviceUrl:  GetEnvString("WEATHER_MICROSERVICE_URL"),
		WeatherMicroserviceName: GetEnvString("WEATHER_MICROSERVICE_NAME"),
	}

	return Config
}

func GetEnvString(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("Environment variable %s is required", key))
	}
	return value
}

func GetEnvNumber(key string) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		panic(fmt.Sprintf("Environment variable %s is required", key))
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		panic(fmt.Sprintf("Environment variable %s must be a valid integer", key))
	}

	return value
}
