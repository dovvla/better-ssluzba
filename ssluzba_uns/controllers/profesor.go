package controllers

import (
	"net/http"
	"ssluzba-uns/models"

	"github.com/gin-gonic/gin"
)

type CreateProfesorInput struct {
	Ime     string `json:"ime" binding:"required"`
	Prezime string `json:"prezime" binding:"required"`
	JMBG    string `json:"jmbg" binding:"required"`
}

func AllProfesors(c *gin.Context) {
	var profesors []models.Profesor
	_ = models.DB.Find(&profesors)
	c.JSON(http.StatusOK, gin.H{"data": profesors})
}

func FindProfesor(c *gin.Context) {
	var profesor models.Profesor
	if err := models.DB.Where("jmbg = ?", c.Param("jmbg")).First(&profesor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": profesor})
}

func CreateProfesor(c *gin.Context) {
	var input CreateProfesorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	profesor := models.Profesor{Ime: input.Ime, Prezime: input.Prezime, JMBG: input.JMBG}
	if err := models.DB.Where("jmbg = ?", input.JMBG).First(&profesor).Error; err != nil {
		models.DB.Create(&profesor)
		c.JSON(http.StatusOK, gin.H{"data": profesor})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Profesor sa datim jmbg-om vec postoji"})
}
