package usecases

import (
	"cep_weather/application/app"
	"cep_weather/application/interfaces"
	"cep_weather/application/usecases/dtos"
	"fmt"
	"net/http"
)

type WeatherUseCase struct {
	weatherRepository interfaces.IWeatherRepository
}

func NewWeatherUseCase(wr interfaces.IWeatherRepository) *WeatherUseCase {
	return &WeatherUseCase{weatherRepository: wr}
}

func (r *WeatherUseCase) SearchByCity(city string) (dtos.WeatherUseCaseOutput, app.Errors) {
	data, err := r.weatherRepository.GetWeather(city)

	if err != nil {
		return dtos.WeatherUseCaseOutput{}, app.CreateErrors(app.Error{
			Code: http.StatusUnprocessableEntity,
			Type: app.ERROR_UNKNOW,
		})
	}

	return dtos.WeatherUseCaseOutput{
		TempC: fmt.Sprintf("%.2f", data.Current.TempC),
		TempF: transformCelsiusToFahrenheit(data.Current.TempC),
		TempK: transformCelsiusToKelvin(data.Current.TempC),
	}, nil
}

func transformCelsiusToFahrenheit(celsius float64) string {
	fahrenheightStr := fmt.Sprintf("%.2f", celsius*1.8+32)

	return fahrenheightStr
}

func transformCelsiusToKelvin(celsius float64) string {
	kelvinStr := fmt.Sprintf("%.2f", celsius+273)

	return kelvinStr
}
