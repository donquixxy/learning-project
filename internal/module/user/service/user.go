package service

import (
	"context"
	"learning-project/config"
	"learning-project/internal/app"
	"learning-project/internal/module/user/entity"
	"learning-project/internal/module/user/interfaces"
	"learning-project/internal/module/user/payload"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository interfaces.UserRepository
	Commons        *app.AppCommons
}

// Create implements interfaces.UserService.
func (u *UserService) Create(ctx context.Context, data payload.UserCreate) (*entity.User, string, error) {
	return u.userRepository.Create(ctx, data, u.Commons.DB)
}

// Get implements interfaces.UserService.
func (u *UserService) Get(ctx context.Context, data payload.UserGet) (*entity.User, string, error) {
	return u.userRepository.Get(ctx, data, u.Commons.DB)
}

// Update implements interfaces.UserService.
func (u *UserService) Update(ctx context.Context, data payload.UserUpdate) (*entity.User, string, error) {

	if data.Password != nil {
		// Generate new hashed password
		newPassword, err := bcrypt.GenerateFromPassword([]byte(*data.Password), 12)

		if err != nil {
			u.Commons.Logger.Errorf("[Update User] - Failed to generate password for user: %v", err)
			return nil, "failed to update user profile", err
		}

		strPW := string(newPassword)
		data.Password = &strPW
	}

	return u.userRepository.Update(ctx, data, u.Commons.DB)
}

func NewUserService(
	userRepo interfaces.UserRepository,
	Commons *app.AppCommons,
) interfaces.UserService {
	return &UserService{
		userRepository: userRepo,
		Commons:        Commons,
	}
}

func (u *UserService) Login(ctx context.Context, data payload.LoginRequest) (*payload.UserLogin, string, error) {
	// Get user
	user, msg, err := u.userRepository.Get(ctx, payload.UserGet{
		Phone: &data.Phone,
	}, u.Commons.DB)

	if err != nil {
		u.Commons.Logger.Errorf("[Login] - Failed to get user from db: %v", err)
		return nil, msg, err
	}

	// Compare user hashed password with user raw password
	if err = bcrypt.CompareHashAndPassword([]byte((user.Password)), []byte(data.Password)); err != nil {
		u.Commons.Logger.Errorf("[Login] - Invalid user password given. User ID :%v. Received string :%v ", user.ID, data.Password)
		return nil, "Invalid password", err
	}

	// Generate user access token
	token, err := u.GenerateToken(user, false)

	if err != nil {
		u.Commons.Logger.Errorf("[Login] - Failed to generate access token %v", err)
		return nil, "Invalid access token", err
	}

	// Generate refresh token
	refreshToken, err := u.GenerateToken(user, true)

	if err != nil {
		u.Commons.Logger.Errorf("[Login] - Failed to generate refresh token %v", err)
		return nil, "Invalid refresh token", err
	}

	// Update user token at database
	_, msg, err = u.userRepository.Update(ctx, payload.UserUpdate{
		ID:         user.ID,
		LoginToken: &token,
	}, u.Commons.DB)

	if err != nil {
		u.Commons.Logger.Errorf("[Login] - Failed to update user token at database %v", err)
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
		ID:   user.ID,
		Name: user.Name,
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
