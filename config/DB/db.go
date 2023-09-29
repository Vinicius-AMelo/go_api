package config

import (
	"fmt"
	studentModel "student_api/models/Student"

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
		Host:     "localhost",
		User:     "vini",
		Password: "vini",
		Database: "go",
		Port:     "5432",
	}
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
