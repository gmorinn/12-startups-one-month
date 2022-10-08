package utils

import (
	"fmt"
	"strings"
)

func CheckPassword(password string, min int, is_special_caracter bool) string {
	if len(password) < min {
		return fmt.Sprintf("Password must be at least %d characters long", min)
	}
	if strings.Contains(password, " ") {
		return "Password must not contain spaces"
	}
	// check if have numbers
	if !strings.ContainsAny(password, "0123456789") {
		return "Password must contain at least one number"
	}
	// check if have special characters
	if is_special_caracter && !strings.ContainsAny(password, "!@#$%^&*()_+") {
		return "Password must contain at least one special character"
	}
	return ""
}
