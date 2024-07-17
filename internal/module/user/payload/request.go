package payload

type (
	UserCreate struct {
		Name       string
		LoginToken string
		Password   string
		Phone      string
		Email      string
	}

	UserUpdate struct {
		ID         int
		Name       *string
		LoginToken *string
		Password   *string
		Phone      *string
		Email      *string
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
