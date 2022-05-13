package modules

import (
	"github.com/Dan6erbond/go-fx-gin-template/internal/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func ProvideLogger(c *config.AppConfig) *zap.Logger {
	return c.Logger
}

var LoggingModule = fx.Options(
	fx.Provide(ProvideLogger),
)
