package common

import "github.com/go-playground/validator"

var (
	validate = validator.New()
)

func Validate(v any) error {
	return validate.Struct(v)
}
