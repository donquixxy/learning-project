package service

import (
	"context"
	"learning-project/config"
	"learning-project/internal/module/user/entity"
	"learning-project/internal/module/user/interfaces"
	"learning-project/internal/module/user/payload"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	userRepository interfaces.UserRepository
	DB             *gorm.DB
}

// Create implements interfaces.UserService.
func (u *UserService) Create(ctx context.Context, data payload.UserCreate) (*entity.User, string, error) {
	return u.userRepository.Create(ctx, data, u.DB)
}

// Get implements interfaces.UserService.
func (u *UserService) Get(ctx context.Context, data payload.UserGet) (*entity.User, string, error) {
	return u.userRepository.Get(ctx, data, u.DB)
}

// Update implements interfaces.UserService.
func (u *UserService) Update(ctx context.Context, data payload.UserUpdate) (*entity.User, string, error) {
	return u.userRepository.Update(ctx, data, u.DB)
}

func NewUserService(
	DB *gorm.DB,
	userRepo interfaces.UserRepository,
) interfaces.UserService {
	return &UserService{
		userRepository: userRepo,
		DB:             DB,
	}
}

func (u *UserService) Login(ctx context.Context, data payload.LoginRequest) (*payload.UserLogin, string, error) {
	// Get user
	user, msg, err := u.userRepository.Get(ctx, payload.UserGet{
		Phone: &data.Phone,
	}, u.DB)

	if err != nil {
		return nil, msg, err
	}

	// Compare user hashed password with user raw password
	if err = bcrypt.CompareHashAndPassword([]byte((user.Password)), []byte(data.Password)); err != nil {
		return nil, "Invalid password", err
	}

	// Generate user access token
	token, err := u.GenerateToken(user, false)

	if err != nil {
		log.Printf("Failed to generate access token: %v", err)
		return nil, "Invalid access token", err
	}

	// Generate refresh token
	refreshToken, err := u.GenerateToken(user, true)

	if err != nil {
		log.Printf("Failed to generate refresh token %v", err)
		return nil, "Invalid refresh token", err
	}

	// Update user token at database
	_, msg, err = u.userRepository.Update(ctx, payload.UserUpdate{
		ID:         user.ID,
		LoginToken: &token,
	}, u.DB)

	if err != nil {
		return nil, msg, err
	}

	result := payload.UserLogin{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}

	return &result, "OK", nil
}

func (u *UserService) GenerateToken(user *entity.User, isRefreshToken bool) (string, error) {
	appConfig := config.GetJwtConfig()

	now := time.Now()
	expDate := time.Now().Add(time.Hour * 72)
	claims := payload.JwtClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    appConfig.Issuer,
			Subject:   "subject",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expDate),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(appConfig.Secret)

	if isRefreshToken {
		secret = []byte(appConfig.SecretRefresh)
	}

	strToken, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return strToken, nil
}
