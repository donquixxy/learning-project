//go:build wireinject
// +build wireinject

package driver

import (
	"learning-project/internal/driver"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var dbWire = wire.NewSet(
	driver.InitDatabase,
)

func InitializeDatabase() *gorm.DB {
	wire.Build(dbWire)

	return &gorm.DB{}
}
