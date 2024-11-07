package env_test

import (
	env "cep_weather/infra/config"
	"os"
	"testing"
)

func setupEnvVars() {
	os.Setenv("WEATHER_API_KEY", "test-api-key")
	os.Setenv("PORT", "8080")
}

func clearEnvVars() {
	os.Unsetenv("WEATHER_API_KEY")
	os.Unsetenv("PORT")
}

func TestLoadEnvs(t *testing.T) {
	setupEnvVars()
	defer clearEnvVars()

	config := env.LoadEnvs()

	if config.WeatherApiKey != "test-api-key" {
		t.Errorf("Expected WEATHER_API_KEY to be 'test-api-key', got %s", config.WeatherApiKey)
	}

	if config.Port != 8080 {
		t.Errorf("Expected PORT to be 8080, got %d", config.Port)
	}
}

func TestGetEnvString_Success(t *testing.T) {
	os.Setenv("TEST_STRING_ENV", "value")
	defer os.Unsetenv("TEST_STRING_ENV")

	value := env.GetEnvString("TEST_STRING_ENV")
	if value != "value" {
		t.Errorf("Expected TEST_STRING_ENV to be 'value', got %s", value)
	}
}

func TestGetEnvString_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected GetEnvString to panic when variable is missing")
		}
	}()

	os.Unsetenv("MISSING_STRING_ENV")
	env.GetEnvString("MISSING_STRING_ENV")
}

func TestGetEnvNumber_Success(t *testing.T) {
	os.Setenv("TEST_NUMBER_ENV", "1234")
	defer os.Unsetenv("TEST_NUMBER_ENV")

	value := env.GetEnvNumber("TEST_NUMBER_ENV")
	if value != 1234 {
		t.Errorf("Expected TEST_NUMBER_ENV to be 1234, got %d", value)
	}
}

func TestGetEnvNumber_PanicOnMissing(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected GetEnvNumber to panic when variable is missing")
		}
	}()

	os.Unsetenv("MISSING_NUMBER_ENV")
	env.GetEnvNumber("MISSING_NUMBER_ENV")
}

func TestGetEnvNumber_PanicOnInvalidValue(t *testing.T) {
	os.Setenv("INVALID_NUMBER_ENV", "invalid")
	defer os.Unsetenv("INVALID_NUMBER_ENV")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected GetEnvNumber to panic on invalid integer value")
		}
	}()

	env.GetEnvNumber("INVALID_NUMBER_ENV")
}

func TestLoadEnvs_ConfigAlreadyInitialized(t *testing.T) {
	setupEnvVars()
	defer clearEnvVars()

	firstConfig := env.LoadEnvs()
	secondConfig := env.LoadEnvs()

	if firstConfig != secondConfig {
		t.Errorf("Expected LoadEnvs to return the same instance, but got different instances")
	}
}
