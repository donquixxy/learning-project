package service

import (
	"context"
	"errors"
	"learning-project/internal/app"
	"learning-project/internal/module/user/entity"
	"learning-project/internal/module/user/interfaces"
	"learning-project/internal/module/user/payload"
	"time"
)

type AttendanceService struct {
	AttendanceRepo interfaces.AttendanceRepository
	Commons        *app.AppCommons
}

func (a AttendanceService) Create(ctx context.Context, data payload.AttendanceCreate) (*entity.Attendances, string, error) {
	if data.ClockInStr != "" {
		parsedClockIn, err := time.Parse("2006-01-02 15:04:05", data.ClockInStr)

		if err != nil {
			a.Commons.Logger.Errorf("[CreateAttendance]  - failed to parse clockin %v. given value %v", err, data.ClockInStr)
			return nil, "invalid clock-in time given", err
		}

		data.ClockinTime = parsedClockIn
	}

	if data.ClockOutStr != "" {
		parsedClockOut, err := time.Parse("2006-01-02", data.ClockOutStr)

		if err != nil {
			return nil, "invalid clock-out time given", err
		}

		data.ClockoutTime = &parsedClockOut
	}

	// Validate wether used already clock-in for today
	attendance, _, _ := a.AttendanceRepo.Get(ctx, payload.AttendanceGet{
		UserID:      &data.UserID,
		ClockInTime: &data.ClockinTime,
	}, a.Commons.DB)

	if attendance != nil {
		return nil, "you already clock-in today", errors.New("you already clock-in today")
	}

	// Send message to queue
	return a.AttendanceRepo.Create(ctx, data, a.Commons.DB)
}

func (a AttendanceService) Update(ctx context.Context, data payload.AttendanceUpdate) (*entity.Attendances, string, error) {
	return a.AttendanceRepo.Update(ctx, data, a.Commons.DB)
}

func (a AttendanceService) Get(ctx context.Context, data payload.AttendanceGet) (*entity.Attendances, string, error) {
	return a.AttendanceRepo.Get(ctx, data, a.Commons.DB)
}

func NewAttendanceService(
	Commons *app.AppCommons,
	AttendanceRepo interfaces.AttendanceRepository,
) interfaces.AttendanceService {
	return &AttendanceService{
		AttendanceRepo: AttendanceRepo,
		Commons:        Commons,
	}
}
