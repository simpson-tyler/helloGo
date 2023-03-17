package connection

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbConnection() (*gorm.DB, error) {
	dbUsername := os.Getenv("postgres_username")
	dbPassword := os.Getenv("postgres_password")
	dbHost := os.Getenv("postgres_host")
	dbPort := os.Getenv("postgres_port")
	dbName := os.Getenv("postgres_db_name")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUsername, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
