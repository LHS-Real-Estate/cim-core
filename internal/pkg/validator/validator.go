package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type CustomValidate struct {
	validate *validator.Validate
}

func NewCustomValidate() *CustomValidate {
	return &CustomValidate{
		validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (cv *CustomValidate) Validate(s interface{}) error {
	err := cv.validate.Struct(s)
	if err != nil {
		var invalidFields []string

		for _, e := range err.(validator.ValidationErrors) {
			invalidFields = append(invalidFields, fmt.Sprintf("%s: \"%s\"", e.Namespace(), e.Value()))
		}

		return fmt.Errorf("invalid fields: %s", strings.Join(invalidFields, ", "))
	}
	return nil
}
