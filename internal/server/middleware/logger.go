package middleware

import (
	"github.com/labstack/echo/v4"
	"learning-project/internal/app"
	"time"
)

func LoggerMiddlware(log *app.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			requestID := c.Response().Header().Get(echo.HeaderXRequestID)
			c.Set("start_time", start.Format("2006-01-02 15:04:05"))

			log.Infof("Incoming Request [%v %v]"+
				" IP:%v"+
				" Request ID :%v", c.Request().Method, c.Request().URL, c.RealIP(), requestID)

			defer func() {
				duration := time.Since(start)
				log.Infof("Request ID:%v. Finished request in %v", requestID, duration)
			}()

			// Call the next handler in the chain
			return next(c)
		}
	}
}
