package usecases

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/kelwynOliveira/Goexpert-Deploy-with-Cloud-Run/internal/entity"
)

type FindByCityNameUseCaseInterface interface {
	GetWeather(city string) (*entity.Forecast, error)
}

type FindByCityNameUseCase struct {
	APIKey string
}

func NewFindByCityNameUseCase(
	apiKey string,
) *FindByCityNameUseCase {
	return &FindByCityNameUseCase{
		APIKey: apiKey,
	}
}

func (uc *FindByCityNameUseCase) GetWeather(city string) (*entity.Forecast, error) {
	var weather entity.Forecast

	request, err := http.Get(
		fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", uc.APIKey, url.QueryEscape(city)),
	)
	if err != nil {
		return nil, err
	}

	result, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, &weather)
	if err != nil {
		return nil, err
	}

	return &weather, err
}
