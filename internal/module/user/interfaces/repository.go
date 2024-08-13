package interfaces

import (
	"context"
	"learning-project/internal/module/user/entity"
	"learning-project/internal/module/user/payload"

	"gorm.io/gorm"
)

// Arguments tx needs to be sent in order to use transaction.
// Otherwise, just pass nil
type UserRepository interface {
	Create(ctx context.Context, data payload.UserCreate, tx *gorm.DB) (*entity.User, string, error)
	Update(ctx context.Context, data payload.UserUpdate, tx *gorm.DB) (*entity.User, string, error)
	Get(ctx context.Context, data payload.UserGet, tx *gorm.DB) (*entity.User, string, error)
}

type AttendanceRepository interface {
	Create(ctx context.Context, data payload.AttendanceCreate, tx *gorm.DB) (*entity.Attendances, string, error)
	Update(ctx context.Context, data payload.AttendanceUpdate, tx *gorm.DB) (*entity.Attendances, string, error)
	Get(ctx context.Context, data payload.AttendanceGet, tx *gorm.DB) (*entity.Attendances, string, error)
}
