package repository

import (
	"context"
	"learning-project/internal/module/user/entity"
	"learning-project/internal/module/user/interfaces"
	"learning-project/internal/module/user/payload"
	"time"

	"gorm.io/gorm"
)

type userRepository struct {
}

// Create implements interfaces.UserRepository.
func (u *userRepository) Create(ctx context.Context, data payload.UserCreate, tx *gorm.DB) (*entity.User, string, error) {
	now := time.Now()
	payload := entity.User{
		Name:       data.Name,
		LoginToken: data.LoginToken,
		Password:   data.Password,
		Phone:      data.Phone,
		Email:      data.Email,
		CreatedAt:  now,
		UpdatedAt:  &now,
	}

	query := tx.WithContext(ctx)

	if err := query.Create(&payload).Error; err != nil {
		return nil, "failed to create user", err
	}

	return &payload, "successfully created user", nil
}

// Get implements interfaces.UserRepository.
func (u *userRepository) Get(ctx context.Context, data payload.UserGet, tx *gorm.DB) (*entity.User, string, error) {
	var result entity.User
	query := tx.WithContext(ctx)

	if data.Email != nil {
		query = query.Where("email = ?", *data.Email)
	}

	if data.ID != nil {
		query = query.Where("id = ?", *data.ID)
	}

	if data.Password != nil {
		query = query.Where("password = ?", *data.Password)
	}

	if data.Phone != nil {
		query = query.Where("phone = ?", *data.Phone)
	}

	if err := query.First(&result).Error; err != nil {
		return nil, "failed to retrieve user", err
	}

	return &result, "successfully retrieved user", nil
}

// Update implements interfaces.UserRepository.
func (u *userRepository) Update(ctx context.Context, data payload.UserUpdate, tx *gorm.DB) (*entity.User, string, error) {
	var user *entity.User

	query := tx.WithContext(ctx)

	if err := query.Where("id = ?", data.ID).First(&user).Error; err != nil {
		return nil, "failed to update user", err
	}

	if data.Name != nil {
		user.Name = *data.Name
	}

	if data.LoginToken != nil {
		user.LoginToken = *data.LoginToken
	}

	if data.Password != nil {
		user.Password = *data.Password
	}

	if data.Email != nil {
		user.Email = *data.Email
	}

	if data.Phone != nil {
		user.Phone = *data.Phone
	}

	if err := tx.WithContext(ctx).Save(&user).Error; err != nil {
		return nil, "failed to save user", err
	}

	return user, "OK", nil
}

func NewUserRepository() interfaces.UserRepository {
	return &userRepository{}
}
