package studentController

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	config "student_api/config/DB"
	studentModel "student_api/models/Student"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var Students []studentModel.Student

func GetStudents(c *gin.Context) {
	db := config.DBConnection()
	var students []studentModel.Student
	result := db.Find(&students)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, students)

}

func GetStudentByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
		})

		return
	}

	db := config.DBConnection()
	var student studentModel.Student
	result := db.First(&student, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Aluno não encontrado",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao encontrar aluno",
			})
		}

		return
	}

	c.JSON(http.StatusOK, student)

}

func PostStudent(c *gin.Context) {
	var student studentModel.Student

	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao processar dados",
			"error":   err.Error(),
		})

		return
	}

	validate := validator.New()

	err := validate.Struct(student)

	if err != nil {
		var missingFields []string
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			missingFields = append(missingFields, strings.ToLower(fieldName))
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"missing_fields": missingFields,
		})

		return
	}

	db := config.DBConnection()
	result := db.Create(&student)
	c.JSON(http.StatusOK, result)

}

func PutStudent(c *gin.Context) {
	var reqStudent studentModel.Student
	idStr := c.Param("id")

	err := c.ShouldBindJSON(&reqStudent)

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
		})

		return
	}

	fmt.Print(id)

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Estudante não encontrado",
	})

}

func DeleteStudent(c *gin.Context) {
	idrStr := c.Param("id")

	id, err := strconv.Atoi(idrStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID Inválido",
		})

		return
	}

	if id > 0 && id < len(Students) {
		Students = append(Students[:id], Students[id+1:]...)
	}

	c.JSON(http.StatusOK, Students)
}
