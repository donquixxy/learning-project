package user

import (
	"learning-project/internal/app"
	"learning-project/internal/module/user/handler"
	"learning-project/internal/module/user/wire"
	"learning-project/internal/server"
)

func InitUserModule(router *server.Router, commons *app.AppCommons) {

	userService := wire.InitUserService(commons)
	attendanceService := wire.InitAttendanceService(commons)
	userHandler := handler.NewUserHandler(userService)
	attendanceHandler := handler.NewAttendanceHandler(attendanceService)
	router.PublicApi.POST("/auth", userHandler.Login)
	router.PrivateApi.PUT("/user", userHandler.Update)
	router.PrivateApi.POST("/attendance", attendanceHandler.Create)
	router.PrivateApi.GET("/attendance", attendanceHandler.Get)
}
