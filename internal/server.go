package internal

import (
	"github.com/Dan6erbond/go-fx-gin-template/internal/config"
	"github.com/Dan6erbond/go-fx-gin-template/internal/controllers"
	"github.com/Dan6erbond/go-fx-gin-template/internal/router"
	"github.com/Dan6erbond/go-fx-gin-template/internal/services"
	"github.com/gin-gonic/gin"
)

func NewGin() *gin.Engine {
	r := gin.Default()
	return r
}

func MakeServer(c ...*config.AppConfig) *gin.Engine {
	appConfig := config.ProvideConfig(c...)

	r := gin.Default()

	helloWorldService := services.NewHelloWorldService(appConfig.Logger)
	helloWorldController := controllers.NewHelloWorldController(helloWorldService)

	router.RegisterRoutes(r, helloWorldController)

	return r
}
