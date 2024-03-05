package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetViaCEP(t *testing.T) {
	result, err := GetViaCEP("01153000")
	assert.Nil(t, err)
	assert.Equal(t, "01153-000", result.Zipcode)
	assert.Equal(t, "Rua Vitorino Carmilo", result.AddressLine1)
	assert.Equal(t, "", result.AddressLine2)
	assert.Equal(t, "Barra Funda", result.Neighborhood)
	assert.Equal(t, "São Paulo", result.City)
	assert.Equal(t, "SP", result.State)
	assert.Equal(t, "3550308", result.IBGE)
	assert.Equal(t, "1004", result.GIA)
	assert.Equal(t, "11", result.Area)
	assert.Equal(t, "7107", result.SIAFI)
}
