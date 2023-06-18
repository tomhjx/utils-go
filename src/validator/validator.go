package validator

import (
	"github.com/go-playground/validator/v10"
)

type Validate struct {
	validator.Validate
}

func New() *Validate {
	v := validator.New()
	v = registerCustomValidations(v)
	return &Validate{
		Validate: *v,
	}
}
