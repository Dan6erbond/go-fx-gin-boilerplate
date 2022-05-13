package modules

import (
	"github.com/Dan6erbond/go-fx-gin-template/internal/router"
	"go.uber.org/fx"
)

var ApiModule = fx.Options(
	fx.Invoke(router.RegisterRoutes),
)
