package modules

import (
	"github.com/Dan6erbond/go-fx-gin-template/internal/controllers"
	"github.com/Dan6erbond/go-fx-gin-template/internal/services"
	"go.uber.org/fx"
)

var HelloWorldModule = fx.Options(
	fx.Provide(services.NewHelloWorldService),
	fx.Provide(controllers.NewHelloWorldController),
)
