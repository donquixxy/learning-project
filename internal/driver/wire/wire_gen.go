// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package driver

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"learning-project/internal/driver"
)

// Injectors from wire.go:

func InitializeDatabase() *gorm.DB {
	db := driver.InitDatabase()
	return db
}

// wire.go:

var dbWire = wire.NewSet(driver.InitDatabase)
