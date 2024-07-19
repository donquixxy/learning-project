package user

import (
	"learning-project/internal/module/user/handler"
	"learning-project/internal/module/user/wire"
	"learning-project/internal/server"
)

func InitUserModule(router *server.Router) {

	service := wire.InitUserService()
	handler := handler.NewUserHandler(service)
	router.PublicApi.POST("/auth", handler.Login)
}
