package utils

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/jhonnydsl/clinify-backend/src/dtos"
)

func ValidateAdminInput(admin dtos.AdminInput) error {
	fullname := strings.TrimSpace(admin.FullName)

	if utf8.RuneCountInString(fullname) < 5 {
		return fmt.Errorf("name must be at least 5 characters long")
	}

	parts := strings.Fields(fullname)
	if len(parts) < 2 {
		return fmt.Errorf("full name must include first and last name")
	}

	_, err := mail.ParseAddress(admin.Email)
	if err != nil {
		return fmt.Errorf("invalid email format")
	}

	if utf8.RuneCountInString(admin.Password) < 6 {
		return fmt.Errorf("the password must be at least 6 characters long")
	}

	birth := admin.BirthDate

	if birth.IsZero() {
		return fmt.Errorf("birth date is required")
	}

	if birth.After(time.Now()) {
		return fmt.Errorf("birth date cannot be in the future")
	}

	age := time.Now().Year() - birth.Year()

	if time.Now().YearDay() < birth.YearDay() {
		age--
	}

	if age < 18 {
		return fmt.Errorf("minimum age is 18 years old")
	}

	phone := strings.TrimSpace(admin.Phone)

	normalized := regexp.MustCompile(`\D`).ReplaceAllString(phone, "")

	if len(normalized) < 10 || len(normalized) > 11 {
		return fmt.Errorf("invalid phone number")
	}

	return nil
}