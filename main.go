package main

import (
	"fmt"
	"os"
	"regexp"
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
		// Ensure the word has at least one character
		if len(words[i]) == 0 {
			continue
		}

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

		// Check if the word contains a "(cap," pattern
		// Use a regular expression to match the "(cap, N)" pattern
		re := regexp.MustCompile(`\(cap,\s*(\d+)\s*\)`)

		// Check if the word matches the pattern
		matches := re.FindStringSubmatch(words[i])
		// fmt.Println("matches:", matches)

		if len(matches) > 0 {
			// Extract the number inside the parentheses
			insideParens := matches[1]

			// Debug: Print the extracted value
			fmt.Println("Extracted from parentheses:", insideParens)

			// Convert the number to an integer
			numWords, err := strconv.Atoi(insideParens)

			/* 	}
			if strings.HasPrefix(words[i], "(cap,") {
				// Extract the number of words to capitalize
				insideParens := strings.TrimSuffix(strings.TrimPrefix(words[i], "(cap,"), ")")
				// Debugging: Print what is extracted from inside the parentheses
				fmt.Println("Extracted from parentheses:", insideParens)
				// Remove any leading/trailing spaces
				insideParens = strings.TrimSpace(insideParens)
				// convert number inside word to int
				numWords, err := strconv.Atoi(insideParens)
				// Debug: Print the extracted number of words
				fmt.Println("numWords:", numWords)
				// // Debug: Show which words are being capitalized
				// fmt.Println("Capitalizing words from index", startIdx, "to", i-1) */

			if err == nil && numWords > 0 {
				// Ensure there's enough words to capitalize
				// Check that we have at least one word before (cap,)
				if i >= 1 {
					// Calculate the start index for capitalization
					fmt.Println("numwords no error:", numWords)
					startIdx := i - numWords
					if startIdx < 0 {
						startIdx = 0 // Adjust to start from the first word if we exceed bounds
					}

					// Debug: Show which words are being capitalized
					fmt.Println("Capitalizing words from index", startIdx, "to", i-1)

					// Capitalize the previous 'numWords' words
					for j := startIdx; j < i; j++ {
						words[j] = textutils.CapitalizeWord(words[j])
					}
				} else {
					fmt.Println("Error: insufficient words to capitalize:", insideParens)
				}
			} else {
				fmt.Println("Error: invalid number for capitalization:", insideParens)
			}
			// Remove the (cap, N) keyword
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

		// Initialize index to track the first non-punctuation character
		charIndex := 0
		// Get the first character of the word at index 0
		// char := string(words[i][charIndex])

		// Count number of leading punctuation characters
		// for charIndex < len(words[i]) {
		// if textutils.IsPunctuation(firstChar) && i > 0 {
		for charIndex < len(words[i]) && textutils.IsPunctuation(string(words[i][charIndex])) {
			charIndex++
		}

		// charIndex indicates if there are leading puntuation characters and a previous word exists i > 0.
		if charIndex > 0 && i > 0 {
			// Extract the leading punctuation characters into punctuation variable by slicing up to charIndex
			punctuation := words[i][:charIndex]
			fmt.Println("Punctuation characters extracted:", punctuation)

			// Add punctuation characters to previous word:
			words[i-1] = words[i-1] + punctuation
			fmt.Println("Moved punctuation", punctuation, "to", words[i-1])

			// If the word has more than one punctuation character, remove
			// if len(words[i]) > 1 {
			// Remove the leading punctuationfrom the current word
			words[i] = words[i][charIndex:] // Slice from the character index onward

			// If more than one punctuation char/groups of punctuation like: ... or !?
			// } else {
			// If the current word becomes empty after removing punctuation, delete it
			if words[i] == "" {
				fmt.Println("one word removed:", words[i])
				words = append(words[:i], words[i+1:]...) // Remove the empty word
				i--                                       // Adjust index since we removed one word
			}
		}

		// The punctuation mark ' will always be found with another instance of it
		// if words[i] == "'" && i > 0 {
		// 	continue
		// }

	}

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
