package utils

import (
	"strings"
)

// FindInSeparated checks whether substr is presented in s as one of separated values
// delimiters is string contains all possible delimiters
func FindInSeparated(s string, substr string, delimiters string) bool {
	if len(substr) == 0 {
		return true
	}

	var (
		runes       = []rune(s)
		runesSubstr = []rune(substr)
	)

	if len(runes) < len(runesSubstr) {
		return false
	}

	var commonPrefixLen int

	for i := 0; i < len(runes); i++ {
		currentRune := runes[i]

		if strings.ContainsRune(delimiters, currentRune) && commonPrefixLen == len(runesSubstr) {
			return true
		}

		if commonPrefixLen < len(runesSubstr) && currentRune == runesSubstr[commonPrefixLen] {
			commonPrefixLen++
			continue
		}

		commonPrefixLen = 0

		for i < len(runes) && !strings.ContainsRune(delimiters, runes[i]) {
			i++
		}
	}

	return commonPrefixLen == len(substr)
}
