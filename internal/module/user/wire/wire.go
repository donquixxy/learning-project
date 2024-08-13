//go:build wireinject
// +build wireinject

package wire

import (
	"learning-project/internal/app"
	"learning-project/internal/module/user/interfaces"
	"learning-project/internal/module/user/repository"
	"learning-project/internal/module/user/service"

	"github.com/google/wire"
)

var userSet = wire.NewSet(
	repository.NewUserRepository,
	service.NewUserService,
)

func InitUserService(commons *app.AppCommons) interfaces.UserService {
	wire.Build(userSet)
	return &service.UserService{}
}

var attendanceSet = wire.NewSet(
	repository.NewAttendanceRepository,
	service.NewAttendanceService,
)

func InitAttendanceService(commons *app.AppCommons) interfaces.AttendanceService {
	wire.Build(attendanceSet)
	return &service.AttendanceService{}
}
