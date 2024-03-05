package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"github.com/kelwynOliveira/Goexpert-Deploy-with-Cloud-Run/internal/entity"
	"github.com/kelwynOliveira/Goexpert-Deploy-with-Cloud-Run/internal/usecases"
)

type WebClimateHandlerInterface interface {
	TemperatureHandler(w http.ResponseWriter, r *http.Request)
}

type WebClimateHandler struct {
	FindClimateByCityNameUseCase usecases.FindByCityNameUseCaseInterface
}

func NewWebClimateHandler(
	findByCityNameUC usecases.FindByCityNameUseCaseInterface,
) *WebClimateHandler {
	return &WebClimateHandler{
		FindClimateByCityNameUseCase: findByCityNameUC,
	}
}

func (h *WebClimateHandler) TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	zipStr := r.URL.Query().Get("zipcode")

	err := validateInput(zipStr)
	if err != nil {
		msgError(err, http.StatusUnprocessableEntity, w)
		return
	}

	location, err := usecases.GetViaCEP(zipStr)
	if err != nil {
		msgError(err, http.StatusInternalServerError, w)
		return
	}
	if location.City == "" {
		err = errors.New("can not find zipcode")
		msgError(err, http.StatusNotFound, w)
		return
	}

	climate, err := h.FindClimateByCityNameUseCase.GetWeather(location.City)
	if err != nil {
		msgError(err, http.StatusInternalServerError, w)
		return
	}

	fahrenheit, kelvin := convertTemperature(climate.Current.TempC)

	response := entity.Temperature{
		Celcius:    float32(climate.Current.TempC),
		Fahrenheit: float32(fahrenheit),
		Kelvin:     float32(kelvin),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func convertTemperature(celcius float64) (float64, float64) {
	fahrenheit := celcius*1.8 + 32
	kelvin := celcius + 273.15

	return fahrenheit, kelvin
}

func validateInput(zipCode string) error {
	matched, err := regexp.MatchString(`\b\d{8}\b`, zipCode)
	if !matched || err != nil {
		return errors.New("invalid zipcode")
	}

	return nil
}

func msgError(err error, statusCode int, w http.ResponseWriter) {
	msg := struct {
		Message string `json:"message"`
	}{
		Message: err.Error(),
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(msg)
}
