package controllers

import (
	"net/http"

	"github.com/Dan6erbond/go-fx-gin-template/internal/dto"
	"github.com/Dan6erbond/go-fx-gin-template/internal/services"
	"github.com/gin-gonic/gin"
)

type HelloWorldController struct {
	jobsService *services.HelloWorldService
}

func NewHelloWorldController(jobsService *services.HelloWorldService) *HelloWorldController {
	return &HelloWorldController{jobsService: jobsService}
}

func (c *HelloWorldController) GetHelloWorld(ctx *gin.Context) {
	var request dto.HelloWorldRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.jobsService.GetHelloWorld(request.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, response)
}
