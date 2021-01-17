package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointService(t *testing.T) {

	t.Run("Should be throw an error if file does not exists", func(t *testing.T) {
		file := "file_does_not_exists.csv"
		testcsv := NewCsvService()
		_, err := testcsv.ParseCSVtoPoint(file, true)
		assert.Error(t, err)
	})

	t.Run("Should ignore records if does not match with struct", func(t *testing.T) {
		file := "../data/puntos_examen_fullstack.csv"
		testcsv := NewCsvService()
		result, err := testcsv.ParseCSVtoPoint(file, false)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, len(result), 500)
	})

	t.Run("Should return a list of endpoints correctly", func(t *testing.T) {
		file := "../data/puntos_examen_fullstack.csv"
		testcsv := NewCsvService()
		result, err := testcsv.ParseCSVtoPoint(file, true)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, 500, len(result))
		record1 := result[0]
		assert.Equal(t, "0101000020E6100000609F8C33BED959C0F5E83676C3893940", record1.TheGeom)
		assert.Equal(t, 1, record1.CartoDBID)
		assert.Equal(t, "Sucursal", record1.Type)
		assert.Equal(t, "25.53813876", record1.Latitude)
		assert.Equal(t, "-103.40223397", record1.Longitude)
		assert.Equal(t, "11A579", record1.Color)
	})
}
