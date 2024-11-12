package env

import (
	"fmt"
	"os"
	"strconv"
)

type ConfigMap struct {
	WeatherApiKey string
	ApiPort       int
}

var Config *ConfigMap

func LoadEnvs() *ConfigMap {
	if Config != nil {
		return Config
	}

	Config = &ConfigMap{
		WeatherApiKey: GetEnvString("WEATHER_API_KEY"),
		ApiPort:       GetEnvNumber("API_PORT"),
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
