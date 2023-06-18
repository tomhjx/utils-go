package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	validate := New()
	s := "test@mail.com"
	errs := validate.Var(s, "email")
	assert.Nil(t, errs)

	s = "123"
	errs = validate.Var(s, "email")
	assert.NotNil(t, errs)
}

func TestCustomGooglePlayURL(t *testing.T) {
	validate := New()
	s := "https://play.google.com/store/apps/details?id=com.scopely.monopolygo"
	errs := validate.Var(s, "custom_google_play_url")
	assert.Nil(t, errs)

	s = "https://play2.google.com/store/apps/details?id=com.scopely.monopolygo"
	errs = validate.Var(s, "custom_google_play_url")
	assert.NotNil(t, errs)
}
