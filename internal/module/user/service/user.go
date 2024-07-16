package service

import (
	"context"
	"learning-project/internal/module/user/entity"
	"learning-project/internal/module/user/interfaces"
	"learning-project/internal/module/user/payload"

	"gorm.io/gorm"
)

type UserService struct {
	userRepository interfaces.UserRepository
	DB             *gorm.DB
}

// Create implements interfaces.UserService.
func (u *UserService) Create(ctx context.Context, data payload.UserCreate) (*entity.User, string, error) {
	return u.userRepository.Create(ctx, data, nil)
}

// Get implements interfaces.UserService.
func (u *UserService) Get(ctx context.Context, data payload.UserGet) (*entity.User, string, error) {
	return u.userRepository.Get(ctx, data)
}

// Update implements interfaces.UserService.
func (u *UserService) Update(ctx context.Context, data payload.UserUpdate) (*entity.User, string, error) {
	return u.userRepository.Update(ctx, data, nil)
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
