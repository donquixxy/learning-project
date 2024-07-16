//go:build wireinject
// +build wireinject

package wire

import (
	"learning-project/internal/driver"
	"learning-project/internal/module/user/interfaces"
	"learning-project/internal/module/user/repository"
	"learning-project/internal/module/user/service"

	"github.com/google/wire"
)

var UserServiceSet = wire.NewSet(
	driver.InitDatabase,
	repository.NewUserRepository,
	service.NewUserService,
)

func InitUserService() interfaces.UserService {
	wire.Build(UserServiceSet)

	return &service.UserService{}
}
