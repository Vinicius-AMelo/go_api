package main

import (
	"fmt"
	"os"
	studentController "student_api/controllers/Student"

	"github.com/gin-gonic/gin"
)

func main() {
	envVarValue := os.Getenv("PORT")

	// Se a variável de ambiente estiver vazia, atribua um valor padrão "5000".
	if envVarValue == "" {
		envVarValue = "5000"
	}

	app := gin.Default()

	app.GET("/students", studentController.GetStudents)
	app.GET("/students/:id", studentController.GetStudentByID)
	app.POST("/students", studentController.PostStudent)
	app.PUT("/students/:id", studentController.PutStudent)
	app.DELETE("/students/:id", studentController.DeleteStudent)

	app.Run(fmt.Sprintf(":%s", envVarValue))
}
