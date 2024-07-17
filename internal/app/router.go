package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Router struct {
	PublicApi  *echo.Group
	PrivateApi *echo.Group
	Echo       *echo.Echo
}

func NewRouter() *Router {
	echo := echo.New()
	echo.Validator = NewValidator(validator.New())
	publicApi := echo.Group("")
	privateApi := echo.Group("api")

	router := &Router{
		PublicApi:  publicApi,
		PrivateApi: privateApi,
		Echo:       echo,
	}

	return router
}
