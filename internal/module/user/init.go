package user

import (
	"learning-project/internal/app"
	"learning-project/internal/module/user/handler"
	"learning-project/internal/module/user/wire"
)

func InitUserModule(router *app.Router) {

	service := wire.InitUserService()
	handler := handler.NewUserHandler(service)
	router.PublicApi.POST("/auth", handler.Login)

}
