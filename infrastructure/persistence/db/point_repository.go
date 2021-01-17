package db

import (
	"github.com/Javlopez/opiapi/domain"
	"gorm.io/gorm"
)

//PointRepository struct
type PointRepository struct {
	DB *gorm.DB
}

func (r *PointRepository) SavePoints(p []domain.Point) error {
	err := r.DB.Create(&p).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *PointRepository) Fetch() ([]domain.Point, error) {
	var points []domain.Point
	result := r.DB.Find(&points)

	if result.Error != nil {
		return points, result.Error
	}

	return points, nil
}
