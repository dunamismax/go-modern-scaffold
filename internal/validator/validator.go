package validator

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// CustomValidator is a custom validator for Echo that uses the validator/v10 package.
type CustomValidator struct {
	validator *validator.Validate
}

// NewCustomValidator creates a new CustomValidator.
func NewCustomValidator(validator *validator.Validate) *CustomValidator {
	return &CustomValidator{validator: validator}
}

// Validate validates the given struct.
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you can return a custom error message
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
