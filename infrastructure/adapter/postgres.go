package adapter

import (
	"fmt"
	"github.com/Javlopez/opiapi/infrastructure/persistence/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

// DatabaseAdapter struct
type DatabaseAdapter struct {
	Postgres    *gorm.DB
	dbPointRepo *db.PointRepository
}

// Connect method
func (dba *DatabaseAdapter) Connect() *DatabaseAdapter {

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbSslMode := os.Getenv("DB_SSL_MODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s  dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSslMode,
	)

	if dba.Postgres == nil {
		dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		dba.Postgres = dbConn
	}
	env := os.Getenv("ENV")
	if env != "production" {
		dba.Postgres.Debug()
	}

	return dba
}

func (dba *DatabaseAdapter) DBPointRepository() *db.PointRepository {
	if dba.dbPointRepo == nil {
		dba.dbPointRepo = &db.PointRepository{
			DB: dba.Connect().Postgres,
		}
	}
	return dba.dbPointRepo
}

func (dba *DatabaseAdapter) DB() *gorm.DB {
	return dba.Connect().Postgres
}
