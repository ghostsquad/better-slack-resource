package slackoff

import "gopkg.in/go-playground/validator.v9"

func InitValidator() *validator.Validate {
	return validator.New()
}

type Validatable interface {
	RegisterValidations(val *validator.Validate)
}
