package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func registerCustomValidations(v *validator.Validate) *validator.Validate {
	v.RegisterValidation("custom_google_play_url", customValidateGooglePlayURL)
	return v
}

func customValidateGooglePlayURL(fl validator.FieldLevel) bool {
	url := fl.Field().String()
	regex := regexp.MustCompile(`^https?://play\.google\.com/store/apps/details\?id=[a-zA-Z0-9_]+`)
	return regex.MatchString(url)
}
