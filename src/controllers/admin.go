package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jhonnydsl/clinify-backend/src/dtos"
	"github.com/jhonnydsl/clinify-backend/src/services"
	"github.com/jhonnydsl/clinify-backend/src/utils"
)

type AdminController struct {
	Service *services.AdminService
}

func (controller *AdminController) CreateAdmin(c *gin.Context) {
	var adminInput dtos.AdminInput

	ctx, cancel := utils.NewDBContext()
	defer cancel()

	err := c.ShouldBindJSON(&adminInput)
	if err != nil {
		c.JSON(utils.GetStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	id, err := controller.Service.CreateAdmin(ctx, adminInput)
	if err != nil {
		c.JSON(utils.GetStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"message": "user admin created",
		"id":		id,
	})
}