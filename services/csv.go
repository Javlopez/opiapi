package services

import (
	"encoding/csv"
	"fmt"
	"github.com/Javlopez/opiapi/domain"
	"os"
	"path/filepath"
	"strconv"
)

const ParentDir = "/.."

type CsvService struct {
	RootDir string
}

func NewCsvService() *CsvService {

	currentDir, _ := os.Getwd()
	rootDir := filepath.Dir(currentDir + ParentDir)

	return &CsvService{
		RootDir: rootDir,
	}
}

func (service *CsvService) ParseCSVtoPoint(fileName string, ignoreHeaders bool) ([]domain.Point, error) {
	var points []domain.Point
	file := fmt.Sprintf("%s/%s", service.RootDir, fileName)
	fileOpen, err := os.Open(file)
	if err != nil {
		return points, err
	}
	defer fileOpen.Close()

	csvReader := csv.NewReader(fileOpen)

	if ignoreHeaders {
		if _, err := csvReader.Read(); err != nil {
			return points, err
		}
	}

	records, err := csvReader.ReadAll()
	if err != nil {
		return points, err
	}

	for _, record := range records {

		cartoDBID, _ := strconv.Atoi(record[1])

		if cartoDBID == 0 {
			continue
		}

		data := domain.Point{
			TheGeom:   record[0],
			CartoDBID: cartoDBID,
			Type:      record[2],
			Latitude:  record[3],
			Longitude: record[4],
			Color:     record[5],
		}
		points = append(points, data)
	}

	return points, nil
}
