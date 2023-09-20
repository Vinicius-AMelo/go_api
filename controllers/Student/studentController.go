package studentController

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	studentModel "student_api/models/Student"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Students []studentModel.Student

func GetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, Students)
}

func GetStudentByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Id Inválido",
		})

		return
	}

	var student studentModel.Student

	for _, s := range Students {
		if s.ID == id {
			student = s
			break
		}
	}

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Aluno não encontrado",
		})

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

	Students = append(Students, student)

	c.JSON(http.StatusOK, student)

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
	}

	for _, student := range Students {
		if student.ID == id {
			student = reqStudent
			fmt.Println("-------", reqStudent)
			fmt.Println("-------", student)
		}
	}

	c.JSON(http.StatusOK, Students)

}
