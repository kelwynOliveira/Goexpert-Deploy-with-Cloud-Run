package usecases

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/kelwynOliveira/Goexpert-Deploy-with-Cloud-Run/internal/entity"
)

func GetViaCEP(zipcode string) (*entity.ViaCEP, error) {
	var location entity.ViaCEP

	viaCEPURL := "https://viacep.com.br/ws/" + zipcode + "/json/"

	request, err := http.Get(viaCEPURL)
	if err != nil {
		return nil, err
	}

	result, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, &location)
	if err != nil {
		return nil, err
	}

	return &location, err
}
