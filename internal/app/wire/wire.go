//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"learning-project/internal/app"
)

var logProvider = wire.NewSet(
	app.NewLogger,
)

func InitLogger() *app.Logger {
	wire.Build(logProvider)
	return &app.Logger{}
}
