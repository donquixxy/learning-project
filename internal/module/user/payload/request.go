package payload

import "time"

type (
	UserCreate struct {
		Name       string `json:"name" form:"name" validate:"required"`
		LoginToken string
		Password   string `json:"password" form:"password" validate:"required"`
		Phone      string `json:"phone" form:"phone" validate:"required"`
		Email      string `json:"email" form:"email" validate:"required"`
	}

	UserUpdate struct {
		ID         int
		Name       *string `json:"name" form:"name"`
		LoginToken *string
		Password   *string `json:"password" form:"password"`
		Phone      *string `json:"phone" form:"phone"`
		Email      *string `json:"email" form:"email"`
	}

	UserGet struct {
		ID       *int
		Phone    *string
		Email    *string
		Password *string
	}
)

type (
	LoginRequest struct {
		Phone    string `json:"phone" form:"phone" validate:"required"`
		Password string `json:"password" form:"password" validate:"required"`
	}
)

type (
	AttendanceCreate struct {
		UserID       int
		ClockinTime  time.Time
		ClockoutTime *time.Time
		ClockInStr   string `form:"clock_in" validate:"required"`
		ClockOutStr  string `form:"clock_out"`
	}

	AttendanceUpdate struct {
		ID           int
		ClockInTime  *time.Time
		ClockoutTime *time.Time
		ClockInStr   string
		ClockOutStr  string
		UserID       *int
	}

	AttendanceGet struct {
		ID           *int
		UserID       *int
		ClockInTime  *time.Time
		ClockoutTime *time.Time
		WithUser     bool
	}
)
