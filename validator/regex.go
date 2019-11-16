package validator

import (
	"regexp"
	"unicode"
)

func ValidPassword(s string) bool {
		var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
		if len(s) >= 8 {
		hasMinLen = true
	}
		for _, char := range s {
		switch {
	case unicode.IsUpper(char):
		hasUpper = true
	case unicode.IsLower(char):
		hasLower = true
	case unicode.IsNumber(char):
		hasNumber = true
	case unicode.IsPunct(char) || unicode.IsSymbol(char):
		hasSpecial = true
	}
	}
		return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
	}

func MatchRegex(s string, pattern string) bool  {
	match, _ := regexp.MatchString(pattern,s)

	return match
}