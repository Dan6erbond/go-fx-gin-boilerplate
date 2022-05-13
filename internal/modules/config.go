package modules

import (
	"github.com/Dan6erbond/go-fx-gin-template/internal/config"
	"go.uber.org/fx"
)

func ProvideAppConfig() *config.AppConfig {
	appConfig := config.ProvideConfig()
	return appConfig
}

var ConfigModule = fx.Options(
	fx.Provide(ProvideAppConfig),
)
