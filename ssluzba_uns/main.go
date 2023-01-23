package main

import (
	"ssluzba-uns/controllers"
	"ssluzba-uns/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/profesor/", controllers.AllProfesors)
	r.GET("/profesor/:jmbg", controllers.FindProfesor)
	r.POST("/profesor", controllers.CreateProfesor)
	r.GET("/student", controllers.AllStudents)
	r.GET("/student/:jmbg", controllers.FindStudent)
	r.POST("/student", controllers.CreateStudent)
	r.Run()
}
