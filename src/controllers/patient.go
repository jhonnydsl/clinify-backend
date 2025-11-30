package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jhonnydsl/clinify-backend/src/dtos"
	"github.com/jhonnydsl/clinify-backend/src/services"
	"github.com/jhonnydsl/clinify-backend/src/utils"
)

type PatientController struct {
	Service *services.PatientService
}

func (controller *PatientController) CreatePatient(c *gin.Context) {
	var patientInput dtos.PatientInput

	ctx, cancel := utils.NewDBContext()
	defer cancel()

	err := c.ShouldBindJSON(&patientInput)
	if err != nil {
		c.JSON(utils.GetStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	id, err := controller.Service.CreatePatient(ctx, patientInput)
	if err != nil {
		c.JSON(utils.GetStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"message": "user patient created",
		"id": 		id,
	})
}