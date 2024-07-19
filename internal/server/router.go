package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"learning-project/internal/app"
	localmiddleware "learning-project/internal/server/middleware"
)

type Router struct {
	PublicApi  *echo.Group
	PrivateApi *echo.Group
	Echo       *echo.Echo
}

func NewRouter(log *app.Logger) *Router {
	echo := echo.New()
	echo.Validator = app.NewValidator(validator.New())

	echo.Use(middleware.RequestID())
	echo.Use(localmiddleware.LoggerMiddlware(log))

	publicApi := echo.Group("")
	privateApi := echo.Group("api")

	router := &Router{
		PublicApi:  publicApi,
		PrivateApi: privateApi,
		Echo:       echo,
	}

	return router
}
