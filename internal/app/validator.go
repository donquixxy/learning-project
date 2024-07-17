package app

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator(validator *validator.Validate) *Validator {
	return &Validator{validator: validator}
}

func (v *Validator) Validate(i interface{}) error {

	err := v.validator.Struct(i)

	if err == nil {
		return nil
	}

	errs := err.(validator.ValidationErrors)
	errMsg := ""

	for _, e := range errs {
		if errMsg != "" {
			errMsg += ", "
		}

		errMsg += e.Field() + " " + e.Tag()
	}

	return errors.New(errMsg)
}
