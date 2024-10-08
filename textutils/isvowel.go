package textutils

import "strings"

// IsVowel checks if a letter is a vowel (a, e, i, o, u)
// This function replaces if:
// if firstLetter == "a" || firstLetter == "e" || firstLetter == "i" || firstLetter == "o" || firstLetter == "u" ||
//
//	firstLetter == "A" || firstLetter == "E" || firstLetter == "I" || firstLetter == "O" || firstLetter == "U" {
func IsVowel(letter string) bool {
	vowels := "aeiouAEIOU"
	return strings.Contains(vowels, letter)
}
