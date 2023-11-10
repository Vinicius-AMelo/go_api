package config

import (
	"fmt"
	studentModel "student_api/models/Student"

	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type connection struct {
	Host     string
	User     string
	Password string
	Database string
	Port     string
}

var ConnectionConfig connection

func init() {
	ConnectionConfig = connection{
		Host:     getEnv("DB_HOST", "localhost"),
		User:     getEnv("DB_USER", "vini"),
		Password: getEnv("DB_PASSWORD", "vini"),
		Database: getEnv("DB_DATABASE", "go"),
		Port:     getEnv("DB_PORT", "5432"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func DBConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", ConnectionConfig.Host, ConnectionConfig.User, ConnectionConfig.Password, ConnectionConfig.Database, ConnectionConfig.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&studentModel.Student{})

	return db
}
