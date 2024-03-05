package usecases

import (
	"testing"

	"github.com/kelwynOliveira/Goexpert-Deploy-with-Cloud-Run/configs"
	"github.com/stretchr/testify/assert"
)

func TestGetWeather(t *testing.T) {
	config, err := configs.LoadConfig("../../cmd/")
	if err != nil {
		panic(err)
	}

	climate := NewFindByCityNameUseCase(config.WeatherApiKey)

	result, err := climate.GetWeather("Manaus")
	assert.Nil(t, err)
	assert.Equal(t, "Manaus", result.Location.Name)
	assert.NotEmpty(t, result.Current.TempC)
}
