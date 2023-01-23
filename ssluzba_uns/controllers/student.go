package controllers

import (
	"net/http"
	"ssluzba-uns/models"

	"github.com/gin-gonic/gin"
)

type CreateStudentInput struct {
	Ime     string `json:"ime" binding:"required"`
	Prezime string `json:"prezime" binding:"required"`
	JMBG    string `json:"jmbg" binding:"required"`
}

func AllStudents(c *gin.Context) {
	var students []models.Student
	_ = models.DB.Find(&students)
	c.JSON(http.StatusOK, gin.H{"data": students})
}

func FindStudent(c *gin.Context) {
	var student models.Student
	if err := models.DB.Where("jmbg = ?", c.Param("jmbg")).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": student})
}

func CreateStudent(c *gin.Context) {
	var input CreateStudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	student := models.Student{Ime: input.Ime, Prezime: input.Prezime, JMBG: input.JMBG}
	if err := models.DB.Where("jmbg = ?", input.JMBG).First(&student).Error; err != nil {
		models.DB.Create(&student)
		c.JSON(http.StatusOK, gin.H{"data": student})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Student sa datim jmbg-om vec postoji"})
}
