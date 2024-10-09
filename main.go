package main

import (
	"fmt"
	"main/textutils"
	"os"
	"strconv"
	"strings"
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
	fmt.Println("Output filename:", outputFile)

	// Read the contents from the input file into a []byte var inputData
	inputData, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:\n", err) // open <filename>: no such file or directory
		os.Exit(1)
	}

	// Convert the []byte inputData to string for easier manipulation
	inputStr := string(inputData)
	fmt.Println("Original content:\n", inputStr) // Print the string of inputData

	// Split the content into lines based on newlines
	lines := strings.Split(inputStr, "\n")

	// Declare the slice to hold modified lines
	var modifiedLines []string

	for _, line := range lines {
		// Trim leading and trailing whitespace
		line = strings.TrimSpace(line)

		// Skip empty lines
		if line == "" {
			modifiedLines = append(modifiedLines, "")
			continue
		}

		// Split each line into words based on spaces
		words := strings.Fields(line)

		// Process each word
		counter := 0
		for i := 0; i < len(words); i++ {
			// Ensure the word has at least one character
			if len(words[i]) == 0 {
				continue
			}

			word := words[i]
			removeKeyword := false

			// For cases that modify the previous word, ensure `i > 0`
			if i > 0 {
				switch word {
				case "(hex)", "(bin)":
					base := 10
					if word == "(hex)" {
						base = 16
					} else {
						base = 2
					}
					// words[i-1] = textutils.ConvertToDecimal(words[i-1], base)
					decimal, err := strconv.ParseInt(words[i-1], base, 64)
					if err != nil {
						fmt.Println("Error converting hex to decimal:", err)
					} else {
						// Replace the (hex) and hex number with the decimal version of binary
						words[i-1] = fmt.Sprintf("%d", decimal)
						removeKeyword = true
					}

				case "(up)":
					//Convert word to UPPERCASE
					words[i-1] = strings.ToUpper(string(words[i-1]))
					removeKeyword = true

				case "(low)":
					// convert to Lowercase
					words[i-1] = strings.ToLower(string(words[i-1]))
					removeKeyword = true

				case "(cap)":
					// Capitalize the word before "(cap)" using capitalizeWord function
					words[i-1] = textutils.CapitalizeWord(words[i-1])
					removeKeyword = true
				}
			}

			// Handle patterns with regex
			// keyword, err := textutils.FindPatterns(words[i]) // Step 1: Extract the keyword from the current word
			if keyword, err := textutils.FindPatterns(words[i]); err == nil && i+1 < len(words) {

				// numWords, err := textutils.ExtractNumber(words[i+1]) // Step 2: Extract the number from the next word
				if numWords, err := textutils.ExtractNumber(words[i+1]); err == nil && numWords > 0 {
					// Check that we have at least one word before to convert
					if i >= 1 {
						// Calculate the start index of word to cap, up, low
						startIdx := i - numWords
						if startIdx < 0 {
							startIdx = 0 // Adjust to start from the first word if we exceed bounds
						}
						// Capitalize the previous 'numWords' words
						for j := startIdx; j < i; j++ {
							switch keyword {
							case "cap":
								words[j] = textutils.CapitalizeWord(words[j])
							case "up":
								words[j] = strings.ToUpper(words[j])
							case "low":
								words[j] = strings.ToLower(words[j])
							}
						}
					} else {
						fmt.Println("Error: insufficient words to capitalize:", numWords)
					}
					// After processing, remove both the keyword and the number
					words = append(words[:i], words[i+2:]...) // Remove current word and next word (numWords)
					i--                                       // Adjust index since we removed two words
				} else {
					fmt.Println("Error: invalid number for capitalization:", err)
				}
			}

			// Replace "a" with "an" if the next word begins with a vowel.
			if (words[i] == "a" || words[i] == "A") && i+1 < len(words) {
				// Get the first letter of the next word
				firstLetter := string(words[i+1][0])
				// Check if the first letter of the next word is a vowel with helper function isVowel
				if textutils.IsVowel(firstLetter) {
					// Replace "a" with "an" and "A" with "An"
					if words[i] == "a" {
						words[i] = "an"
					} else if words[i] == "A" {
						words[i] = "An"
					}
				}
			}

			// Move puncutation characters and groups: ... or !?
			// Initialize index to track the first non-punctuation character
			charIndex := 0
			for charIndex < len(words[i]) && textutils.IsPunctuation(string(words[i][charIndex])) {
				charIndex++
			}
			// charIndex indicates if there are leading puntuation characters and a previous word exists i > 0.
			if charIndex > 0 && i > 0 {
				// Extract the leading punctuation characters into punctuation variable by slicing up to charIndex
				punctuation := words[i][:charIndex]

				// Add punctuation characters to previous word:
				words[i-1] += punctuation

				// Remove the leading punctuation characters from the current word
				words[i] = words[i][charIndex:] // Slice from the character index onward

				// If the current word becomes empty after removing punctuation, delete it
				if words[i] == "" {
					removeKeyword = true
				}
			}

			// The punctuation mark ' will always be found with another instance of it
			if words[i] == "'" {
				counter++
				// If this is the first instance of the punctuation
				if counter == 1 {
					// Move the punctuation to the start of the next word
					words[i+1] = words[i] + words[i+1]
					removeKeyword = true
					// If this is the second instance of the punctuation
				} else if counter == 2 {
					// Move the punctuation to the end of the previous word
					words[i-1] = words[i-1] + words[i]
					removeKeyword = true
					counter = 0 // Reset counter after handling both instances
				}
			}

			// If a keyword was used, remove it
			if removeKeyword {
				// Remove the keyword from the slice if necessary
				words = append(words[:i], words[i+1:]...)
				i-- // Adjust index since we removed one word
			}
		}

		// Join the modified words back into a single line
		modifiedLine := strings.Join(words, " ")
		modifiedLines = append(modifiedLines, modifiedLine)
	}

	// Join all lines with new lines
	modifiedStr := strings.Join(modifiedLines, "\n")

	// Write the modified contents to the output file, 0644 6=rw Owner, 4=r Group, Others
	err = os.WriteFile(outputFile, []byte(modifiedStr), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(1)
	}

	// Print the modified content to the terminal
	fmt.Println("Modified content:\n", strings.TrimSpace(modifiedStr))
	// fmt.Println(modifiedStr)
}
