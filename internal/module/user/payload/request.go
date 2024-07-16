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
