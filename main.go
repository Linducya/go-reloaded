package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"learn.01founders.co/git/lthomson/go-reloaded/textutils"
)

func main() {
	// Check if two command-line arguments are provided (input and output file names)
	if len(os.Args) != 3 {
		fmt.Println("Text modifier usage: please enter <input file> <output file>")
		os.Exit(1) // Prints exit status 1. The program did not run successfully as the required arguments were not provided.
	}

	// Get the input & output file names from command-line arguemts
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	fmt.Println("Input filename: ", inputFile)
	fmt.Println("Output filename: ", outputFile)

	// Read the contents from the input file into a []byte var inputData
	inputData, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err) // open <filename>: no such file or directory
		os.Exit(1)
	}

	// Modify the contents

	// Convert the []byte inputData to string for easier manipulation
	inputStr := string(inputData)
	// fmt.Println(inputData) // Prints the data in the array of []byte taken from the input file
	fmt.Println("Original Content:")
	// fmt.Println(string(inputData)) // Print the string of inputData
	fmt.Println(inputStr) // Print the string of inputData

	// Split the content into words based on spaces
	words := strings.Fields(inputStr)

	// Loop through words
	for i := 0; i < len(words); i++ {

		// if (hex) should replace the word before with the decimal version of hex
		if words[i] == "(hex)" && i > 0 {
			// Convert the word before (hex) to decimal with hexToDecimal function
			// words[i-1] = strconv.ParseInt(string(words[i-1]), 16, 64)
			// decimal, err := hexToDecimal(string(words[i-1]))
			decimal, err := strconv.ParseInt(words[i-1], 16, 64)
			if err != nil {
				fmt.Println("Error converting hex to decimal:", err)
			} else {
				// Replace the (hex) and hex number with the decimal version of binary
				words[i-1] = fmt.Sprintf("%d", decimal)
				// Remove the (hex) keyword
				words = append(words[:i], words[i+1:]...)
			}
		}

		// Every instance of (bin) should replace the word before with the decimal version of the word
		// (in this case the word will always be a binary number).
		if words[i] == "(bin)" && i > 0 {
			decimal, err := strconv.ParseInt(words[i-1], 2, 64)
			if err != nil {
				fmt.Println("Error converting binary to decimal:", err)
			} else {
				// Replace the (bin) and binary number with the decimal version of binary
				words[i-1] = fmt.Sprintf("%d", decimal)
				// Remove the (bin) keyword
				words = append(words[:i], words[i+1:]...)
				i-- // Adjust index since we removed one word
			}
		}

		// Every instance of (up) converts the word before with the Uppercase version of it.
		if words[i] == "(up)" && i > 0 {
			//Convert word to UPPERCASE
			words[i-1] = strings.ToUpper(string(words[i-1]))
			// Remove the (up) keyword
			words = append(words[:i], words[i+1:]...)
			i-- // Adjust index since we removed one word
		}

		// Every instance of (low) converts the word before with the Lowercase version of it.
		if words[i] == "(low)" && i > 0 {
			// convert to Lowercase
			words[i-1] = strings.ToLower(string(words[i-1]))
			// Remove the (low) keyword
			words = append(words[:i], words[i+1:]...)
			i-- // Adjust index since we removed one word
		}

		// Every instance of (cap) converts the word before with the capitalized version of it.
		if words[i] == "(cap)" && i > 0 {
			//Capitalize the word before "(cap)" using capitalizeWord function
			words[i-1] = textutils.CapitalizeWord(words[i-1])
			// Remove the (cap) keyword
			words = append(words[:i], words[i+1:]...)
			i-- // Adjust index since we removed one word
		}

		// Every instance of a should be turned into an if the next word begins with a vowel (a, e, i, o, u) or a h.
		if (words[i] == "a" || words[i] == "A") && i+1 < len(words) {
			// Get the first letter of the next word
			firstLetter := string(words[i+1][0])
			// Check if the first letter of the next word is a vowel with helper function isVowel
			// if firstLetter == "a" || firstLetter == "e" || firstLetter == "i" || firstLetter == "o" || firstLetter == "u" ||
			// 	firstLetter == "A" || firstLetter == "E" || firstLetter == "I" || firstLetter == "O" || firstLetter == "U" {
			if textutils.IsVowel(firstLetter) {
				// Replace "a" with "an" and "A" with "An"
				if words[i] == "a" {
					words[i] = "an"
				} else if words[i] == "A" {
					words[i] = "An"
				}
			}
		}

		/* 	// Every instance of the punctuations ., ,, !, ?, : and ;
		// should be close to the previous word and with space apart from the next one.
		if textutils.IsPunctuation(words[i]) && i > 0 {
			fmt.Println("Punctuation:", words[i])
			// Next two lines are an alternative:
			// words[i] = words[i-1] + words[i]
			// words = append(words[:i-1], words[i:]...)
			words[i-1] = words[i-1] + words[i]
			fmt.Println("Punctuation:", words[i])
			// words = append(words[:i], words[i-1:]...)
			words = append(words[:i], words[i+1:]...)
			i-- // Adjust index since we removed one word
		} */

		// If case punction is attached to next word, e.g. move comma to previous word
		// Get the first character of the word
		charIndex := 0
		firstChar := string(words[i][charIndex])
		for charIndex < len(words[i]); charIndex++ {
		if textutils.IsPunctuation(firstChar) && i > 0 {
			// add punctuation character to previous word:
			words[i-1] = words[i-1] + firstChar
			fmt.Println("moved punctuation", firstChar, "to ", words[i-1])

			if len(words[i]) > 1 {
					// Remove the first character from the current word
				words[i] = words[i][1:] // Slice from the second character onward

				// If more than one punctuation char/groups of punctuation like: ... or !?

			} else {
				// words[i] = "" // If the word only had one character, set it to an empty string
				words = append(words[:i], words[i+1:]...)
				i-- // Adjust index since we removed one word

			}

			fmt.Println("removed first char punctuation from word[i]", words[i])
		}

		// The punctuation mark ' will always be found with another instance of it

	}

	// Convert all words to UPPERCASE with strings.ToUpper():
	// modifiedData := strings.ToUpper(string(inputData))
	// Convert all words to uppercase with for loop
	// for i := 0; i < len(inputData); i++ {
	// 	if inputData[i] >= 'a' && inputData[i] <= 'z' {
	// 		inputData[i] -= 'a' - 'A'
	// 	}
	// }
	// Print the modifed data to the terminal
	// modifiedData := inputData
	// fmt.Println(string(modifiedData))

	// Join the modifed words back into a single string
	modifiedStr := strings.Join(words, " ")

	// Write the modified contents to the output file, 0644 6=rw Owner, 4=r Group, Others
	err = os.WriteFile(outputFile, []byte(modifiedStr), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(1)
	}

	// Print the modified content to the terminal
	fmt.Println("Modified content:")
	fmt.Println(modifiedStr)

}

// FUNCTIONS

// Helper function to capitalize a word
/* func capitalizeWord(word string) string {
	// Capitalize the alaphabetic part of the word
	// But keep punctuation intact, so we use unicode to handle it properly
	for i, r := range word {
		if unicode.IsLetter(r) {
			// Capitalize the first alphabetic character
			return strings.ToUpper(string(r)) + word[i+1:]
		}
	}
	return word
} */

// Helper function to capitalize just the first letter of a word
/* func capitalizeFirstLetter(word string) string {
	if len(word) == 0 {
		return word
	}
	// Capitalize the first letter and leave the rest of the word unchanged
	runes := []rune(word)
	if unicode.IsLetter(runes[0]) {
		runes[0] = unicode.ToUpper(runes[0])
	}
	return string(runes)
} */

/* // Helper function to convert hex to decimal
func hexToDecimal(hex string) (int64, error) {
	// Convert the hex strting to decimal (base 10)
	decimal, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return 0, err
	}
	return decimal, nil
}
*/

/* // Helper function to check for vowel
  func isVowel(letter string) bool {
	vowels := "aeiouAEIOU"
	return strings.Contains(vowels, letter)
	}
*/

/* // Helper function to remove keyword
func remKeyword(keyword string) {
	words[i-1] = fmt.Sprintf("%d", decimal)
				// Remove the (bin) keyword
				words = append(words[:i], words[i+1:]...)
				i-- // Adjust index since we removed one word

} */

// // Convert the string to slice of byte
// outputData := []byte(modifiedData)
// fmt.Println(string(outputData))

// Variables for keywords
// hex := "(hex)"
// binary := "(bin)"
// uppercase := "(up)"
// lowercase := "(low)"
// caps := "(cap)"
// capNumber := 0
