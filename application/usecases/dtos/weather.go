package dtos

type WeatherUseCaseOutput struct {
	TempC string `json:"temp_C"`
	TempF string `json:"temp_F"`
	TempK string `json:"temp_K"`
}
