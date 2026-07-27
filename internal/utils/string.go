package utils

import (
	"regexp"
)

var alphaASCIIRegex = regexp.MustCompile(`^[a-zA-Z]+$`)

// IsAlphaASCII checks if a string contains only ASCII alphabetic characters.
func IsAlphaASCII(s string) bool {
	return alphaASCIIRegex.MatchString(s)
}
