package models

import "unicode"

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	// Convert the first rune to uppercase and append the rest of the string
	return string(unicode.ToUpper(rune(s[0]))) + s[1:]
}
