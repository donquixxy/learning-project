package server

import (
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
	echo.Validator = app.NewValidator()

	echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	echo.Use(middleware.Recover())
	echo.Use(middleware.RequestID())
	echo.Use(localmiddleware.LoggerMiddlware(log))

	publicApi := echo.Group("")
	privateApi := echo.Group("api")

	privateApi.Use(localmiddleware.ValidateJWT())
	router := &Router{
		PublicApi:  publicApi,
		PrivateApi: privateApi,
		Echo:       echo,
	}

	return router
}
