package textutils

import "strings"

// IsPunctuation checks for instance of punctuation
func IsPunctuation(char string) bool {
	punctuations := ".,!?:;"
	return strings.Contains(punctuations, char)
}
