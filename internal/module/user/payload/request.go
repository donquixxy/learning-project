package payload

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
