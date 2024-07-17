package payload

import "github.com/golang-jwt/jwt/v4"

type (
	UserLogin struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	JwtClaims struct {
		ID int `json:"id"`
		jwt.RegisteredClaims
	}
)
