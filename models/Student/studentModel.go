package studentModel

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name      string    `json:"name" validate:"required"`
	Age       int       `json:"age" validate:"required"`
	Course    string    `json:"course" validate:"required"`
	CreatedAt time.Time `gorm:"-"`
}
