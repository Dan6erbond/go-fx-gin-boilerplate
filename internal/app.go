package internal

import (
	"github.com/Dan6erbond/go-fx-gin-template/internal/modules"
	"go.uber.org/fx"
)

func NewFxApp() *fx.App {
	app := fx.New(
		fx.Provide(NewGin),
		modules.AppModule,
		modules.HelloWorldModule,
	)

	return app
}
