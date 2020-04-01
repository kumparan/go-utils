package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RegexEmail(t *testing.T) {
	t.Run("should return true for valid email", func(t *testing.T) {
		validEmail := "user@domain.com"
		assert.True(t, RegexEmail().MatchString(validEmail))
	})

	t.Run("should return false for invalid email", func(t *testing.T) {
		inValidEmail := []string{
			"user-domain.com",
			"user-domaincom",
			"@user-domain.com",
		}
		for _, email := range inValidEmail {
			assert.False(t, RegexEmail().MatchString(email))
		}
	})
}

func Test_RegexAplhaNumSpace(t *testing.T) {
	tests := []struct {
		Name        string
		Word        string
		ExpectMatch bool
	}{
		{
			Name:        "alphabetical character",
			Word:        "This Is Name",
			ExpectMatch: true,
		},
		{
			Name:        "numerical character",
			Word:        "123",
			ExpectMatch: true,
		},
		{
			Name:        "alphanum character",
			Word:        "Joni123",
			ExpectMatch: true,
		},
		{
			Name:        "alphanumspace character",
			Word:        "Joni 123",
			ExpectMatch: true,
		},
		{
			Name:        "except alphanumspace",
			Word:        "whatzittooya@gmail.com",
			ExpectMatch: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			alphaNumSpace := RegexAlphaNumSpace()
			assert.Equal(t, test.ExpectMatch, alphaNumSpace.MatchString(test.Word))
		})
	}
}
