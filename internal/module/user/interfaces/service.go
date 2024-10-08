package interfaces

import (
	"context"
	"learning-project/internal/module/user/entity"
	"learning-project/internal/module/user/payload"
)

type UserService interface {
	Create(ctx context.Context, data payload.UserCreate) (*entity.User, string, error)
	Update(ctx context.Context, data payload.UserUpdate) (*entity.User, string, error)
	Get(ctx context.Context, data payload.UserGet) (*entity.User, string, error)
	Login(ctx context.Context, data payload.LoginRequest) (*payload.UserLogin, string, error)
}

type AttendanceService interface {
	Create(ctx context.Context, data payload.AttendanceCreate) (*entity.Attendances, string, error)
	Update(ctx context.Context, data payload.AttendanceUpdate) (*entity.Attendances, string, error)
	Get(ctx context.Context, data payload.AttendanceGet) (*entity.Attendances, string, error)
}
