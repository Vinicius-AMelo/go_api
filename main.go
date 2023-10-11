package main

import (
	studentController "student_api/controllers/Student"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/students", studentController.GetStudents)
	app.GET("/students/:id", studentController.GetStudentByID)
	app.POST("/students", studentController.PostStudent)
	app.PUT("/students/:id", studentController.PutStudent)
	app.DELETE("/students/:id", studentController.DeleteStudent)

	app.Run(":5000")
}
