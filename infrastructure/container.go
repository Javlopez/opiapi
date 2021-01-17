package infrastructure

import (
	"github.com/Javlopez/opiapi/infrastructure/adapter"
	"github.com/Javlopez/opiapi/services"
	"gorm.io/gorm"
)

// ContainerInfrastructure interface
type ContainerInfrastructure interface {
	PointService() *services.PointService
	CsvService() *services.CsvService
}

// Container struct
type Container struct {
	adapter.DatabaseAdapter
	pointService *services.PointService
	csvService   *services.CsvService
}

func (c *Container) PointService() *services.PointService {
	if c.pointService == nil {
		c.pointService = &services.PointService{
			PointRepo: c.DBPointRepository(),
		}
	}
	return c.pointService
}

func (c *Container) DBConnection() *gorm.DB {
	return c.DB()
}

func (c *Container) CsvService() *services.CsvService {
	if c.csvService == nil {
		c.csvService = services.NewCsvService()
	}
	return c.csvService
}
