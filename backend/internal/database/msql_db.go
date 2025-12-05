package database

import (
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var ACPS_DB *gorm.DB

func ConnectToMSQLServer() (*gorm.DB, error) {
	// Get connection details from environment variables
	user := os.Getenv("DB_MSQL_ACPS_USER")
	password := os.Getenv("DB_MSQL_ACPS_PASSWORD")
	host := os.Getenv("DB_MSQL_ACPS_HOST")
	port := os.Getenv("DB_MSQL_ACPS_PORT")
	database := os.Getenv("DB_MSQL_ACPS_NAME")

	// Construct the DSN
	dsn := "sqlserver://" + user + ":" + password + "@" + host + ":" + port + "?database=" + database

	// Connect to the database
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
