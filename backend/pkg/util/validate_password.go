package util

import (
	"errors"
	"regexp"
)

func ValidatePassword(password string) error {
	if len(password) < 12 {
		return errors.New("password must be at least 12 characters long")
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return errors.New("password must contain at least one number")
	}
	if !regexp.MustCompile(`[~!@#$%^&*()_+\-={}|:"<>?]`).MatchString(password) {
		return errors.New("password must contain at least one special character (~!@#$%^&*()_+`{}|:\"<>?)")
	}
	return nil
}
