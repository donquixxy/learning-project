//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"learning-project/internal/app"
	"learning-project/internal/driver"
)

var appWireSet = wire.NewSet(
	app.NewLogger,
	app.NewAppCommons,
	app.NewValidator,
	driver.InitDatabase,
	driver.NewRabbitMQConnection,
)

func InitApp() *app.AppCommons {
	wire.Build(appWireSet)

	return &app.AppCommons{}
}
