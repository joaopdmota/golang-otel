package env

import (
	"fmt"
	"os"
	"strconv"
)

type ConfigMap struct {
	APIKey string
	Port   int
}

var Config *ConfigMap

func LoadEnvs() (*ConfigMap, error) {
	if Config != nil {
		return Config, nil
	}

	Config = &ConfigMap{
		APIKey: GetEnvString("API_KEY"),
		Port:   GetEnvNumber("PORT"),
	}

	return Config, nil
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
