package repository

import (
	"context"
	"gorm.io/gorm"
	"learning-project/internal/module/user/entity"
	"learning-project/internal/module/user/interfaces"
	"learning-project/internal/module/user/payload"
	"time"
)

type AttendancesRepository struct {
}

func (a AttendancesRepository) Create(ctx context.Context, data payload.AttendanceCreate, tx *gorm.DB) (*entity.Attendances, string, error) {
	now := time.Now()

	query := tx.WithContext(ctx)

	payload := entity.Attendances{
		UserID:      data.UserID,
		ClockInTime: data.ClockinTime,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	if data.ClockoutTime != nil {
		payload.ClockOutTime = data.ClockoutTime
	}

	if err := query.Create(&payload).Error; err != nil {
		return nil, "failed to create attendances", err
	}

	query.Preload("User")

	return &payload, "successfully create attendance user", nil
}

func (a AttendancesRepository) Update(ctx context.Context, data payload.AttendanceUpdate, tx *gorm.DB) (*entity.Attendances, string, error) {
	var attendances *entity.Attendances

	query := tx.WithContext(ctx)

	if err := query.Where("id = ?", data.UserID).First(&attendances).Error; err != nil {
		return nil, "failed to update attendance", err
	}

	if data.ClockInTime != nil {
		attendances.ClockInTime = *data.ClockInTime
	}

	if data.ClockoutTime != nil {
		attendances.ClockOutTime = data.ClockoutTime
	}

	if data.UserID != nil {
		attendances.UserID = *data.UserID
	}

	if err := query.Save(&attendances).Error; err != nil {
		return nil, "failed to update attendance", err
	}

	return attendances, "successfully update attendance", nil
}

func (a AttendancesRepository) Get(ctx context.Context, data payload.AttendanceGet, tx *gorm.DB) (*entity.Attendances, string, error) {
	query := tx.WithContext(ctx)

	var attendances *entity.Attendances

	if data.ID != nil {
		query = query.Where("id = ?", *data.ID)
	}

	if data.UserID != nil {
		query = query.Where("user_id = ?", *data.UserID)
	}

	if data.ClockoutTime != nil {
		query = query.Where("clock_out_time = ?", *data.ClockoutTime)
	}

	if data.ClockInTime != nil {
		query = query.Where("clock_in_time = ?", *data.ClockInTime)
	}

	if data.WithUser {
		query = query.Preload("User")
	}

	if err := query.First(&attendances).Error; err != nil {
		return nil, "failed to get attendance", err
	}

	return attendances, "successfully get attendance", nil
}

func NewAttendanceRepository() interfaces.AttendanceRepository {
	return &AttendancesRepository{}
}
