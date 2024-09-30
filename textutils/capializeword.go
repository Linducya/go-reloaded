package textutils

import (
	"strings"
	"unicode"
)

// Helper function to capitalize a word
func CapitalizeWord(word string) string {
	// Capitalize the alaphabetic part of the word
	// To keep punctuation intact, use unicode to handle it properly
	for i, r := range word {
		if unicode.IsLetter(r) {
			// Capitalize the first alphabetic character
			return strings.ToUpper(string(r)) + word[i+1:]
		}
	}
	return word
}
