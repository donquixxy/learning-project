//go:build wireinject
// +build wireinject

package wire

import (
	wireapp "learning-project/internal/app/wire"
	wiredb "learning-project/internal/driver/wire"
	"learning-project/internal/module/user/interfaces"
	"learning-project/internal/module/user/repository"
	"learning-project/internal/module/user/service"

	"github.com/google/wire"
)

var userSet = wire.NewSet(
	wiredb.InitializeDatabase,
	repository.NewUserRepository,
	service.NewUserService,
	wireapp.InitLogger,
)

func InitUserService() interfaces.UserService {
	wire.Build(userSet)
	return &service.UserService{}
}
