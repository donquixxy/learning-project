package service

import (
	usermodule "learning-project/internal/module/user/interfaces"
)

type companyService struct {
	UserService usermodule.UserService
}
