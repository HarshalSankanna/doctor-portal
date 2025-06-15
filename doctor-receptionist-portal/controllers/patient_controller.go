package controllers

import (
	"net/http"

	"github.com/HarshalSankanna/doctor-portal/doctor-receptionist-portal/config"
	"github.com/HarshalSankanna/doctor-portal/doctor-receptionist-portal/models"
	"github.com/gin-gonic/gin"
)

func CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&patient)
	c.JSON(http.StatusCreated, patient)
}

func GetPatients(c *gin.Context) {
	var patients []models.Patient
	config.DB.Find(&patients)
	c.JSON(http.StatusOK, patients)
}

func UpdatePatient(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient
	if err := config.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	c.BindJSON(&patient)
	config.DB.Save(&patient)
	c.JSON(http.StatusOK, patient)
}
