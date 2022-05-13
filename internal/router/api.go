package router

import (
	"github.com/Dan6erbond/go-fx-gin-template/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, jobsController *controllers.HelloWorldController) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/hello-world", jobsController.GetHelloWorld)
	}
}
