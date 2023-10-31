package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/vinisioux/go-api/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() (*gorm.DB, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	dbUrl := os.Getenv("DATABASE_URL")

	logger := GetLogger("postgres")

	logger.Info("Initializing Postgres connection")

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		logger.Errorf("Error initializing Postgres: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("Error migrating Postgres: %v", err)
		return nil, err
	}

	logger.Info("Postgres initialized")

	return db, nil
}
