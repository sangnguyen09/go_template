package validator

import (
	"regexp"
	"unicode"
)

func ValidPassword(s string) bool {
	match, _ := regexp.MatchString(`^[a-zA-Z0-9\!\@\#\$\%\^\&\*\(\)\_\+\.\,\;\:]+$`,s)
	if match == false {
			return false
	}
		var (
		hasMinMax  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
		if 8 <= len(s) && len(s) <= 30 {
			hasMinMax = true
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
		return hasMinMax && hasUpper && hasLower && hasNumber && hasSpecial
	}

func MatchRegex(s string, pattern string) bool  {
	match, _ := regexp.MatchString(pattern,s)

	return match
}