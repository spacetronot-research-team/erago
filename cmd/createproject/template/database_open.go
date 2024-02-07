package template

func GetDatabaseOpenTemplate() string {
	return databaseOpenTemplate
}

var databaseOpenTemplate = `package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	sslMode := os.Getenv("SSL_MODE")
	tz := os.Getenv("TZ")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPass, dbName, dbPort, sslMode, tz,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("fail initialize db session: %v", err)
	}

	return db, err
}
`
