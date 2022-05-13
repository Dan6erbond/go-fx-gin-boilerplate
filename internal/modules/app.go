package modules

import "go.uber.org/fx"

var AppModule = fx.Options(
	ConfigModule,
	LoggingModule,
	ApiModule,
)
