package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"learning-project/config"
	user_pld "learning-project/internal/module/user/payload"
	"net/http"
	"strings"
)

func ValidateJWT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cfg := config.GetJwtConfig()
			auth := c.Request().Header.Get("Authorization")

			if auth == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Bearer Token is required")
			}

			splittedTokens := strings.Split(auth, " ")
			bearerPart := splittedTokens[0]
			if len(splittedTokens) != 2 || strings.ToLower(bearerPart) != "bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Bearer Token is invalid")
			}

			userToken := splittedTokens[1]
			claimsUser := user_pld.JwtClaims{}
			claims, err := jwt.ParseWithClaims(userToken, &claimsUser, func(token *jwt.Token) (interface{}, error) {
				return []byte(cfg.Secret), nil
			})

			if !claims.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token is invalid")
			}

			if errors.Is(err, jwt.ErrSignatureInvalid) {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token is invalid")
			}

			if errors.Is(err, jwt.ErrTokenExpired) {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token is expired")
			}

			c.Set("user", &claimsUser)
			return next(c)
		}
	}
}

func GetCurrentUser(c echo.Context) *user_pld.JwtClaims {
	user, ok := c.Get("user").(*user_pld.JwtClaims)

	if !ok {
		return new(user_pld.JwtClaims)
	}

	return user
}
