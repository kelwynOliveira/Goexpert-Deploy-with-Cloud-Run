package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kelwynOliveira/Goexpert-Deploy-with-Cloud-Run/configs"
	"github.com/kelwynOliveira/Goexpert-Deploy-with-Cloud-Run/internal/infra/webserver/handlers"
	"github.com/kelwynOliveira/Goexpert-Deploy-with-Cloud-Run/internal/usecases"
)

func main() {
	config, err := configs.LoadConfig("./")
	if err != nil {
		panic(err)
	}

	climate := usecases.NewFindByCityNameUseCase(config.WeatherApiKey)
	temperatureHandler := handlers.NewWebClimateHandler(climate)

	http.HandleFunc("/", temperatureHandler.TemperatureHandler)
	fmt.Printf("Starting server at port 8080\n")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
