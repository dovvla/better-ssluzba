package main

import (
	"net/http"
	"ssluzba-uns/controllers"
	"ssluzba-uns/models"

	"github.com/gin-gonic/gin"
)

func Increment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.Request
		models.DB.Create(&req)
	}
}

func main() {
	r := gin.Default()
	r.Use(Increment())

	models.ConnectDatabase()

	r.GET("/reqcount/",
		func(c *gin.Context) {
			var requests []models.Request
			_ = models.DB.Find(&requests)
			c.JSON(http.StatusOK, gin.H{"requests": requests})
		})
	r.GET("/profesor/", controllers.AllProfesors)
	r.GET("/profesor/:jmbg", controllers.FindProfesor)
	r.POST("/profesor", controllers.CreateProfesor)
	r.GET("/student", controllers.AllStudents)
	r.GET("/student/:jmbg", controllers.FindStudent)
	r.POST("/student", controllers.CreateStudent)
	r.Run()
}
