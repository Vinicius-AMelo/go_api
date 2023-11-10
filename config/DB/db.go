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
		Host:     getEnv("PGHOST", "localhost"),
		User:     getEnv("PGUSER", "vini"),
		Password: getEnv("PGPASSWORD", "vini"),
		Database: getEnv("PGDATABASE", "go"),
		Port:     getEnv("PGPORT", "5432"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		fmt.Printf("Variável de ambiente %s encontrada. Valor: %s\n", key, value)
		return value
	}
	fmt.Printf("Variável de ambiente %s não encontrada. Usando valor padrão: %s\n", key, fallback)
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
