package textutils

import (
	"regexp"
	"strconv"
)

// Helper function to extract number from next word
func ExtractNumber(nextWord string) (int, error) {
	// Compile a regex pattern to capture the number
	numberPattern := `(\d+)\)`
	re := regexp.MustCompile(numberPattern)

	// Find the match and capture the number
	matches := re.FindStringSubmatch(nextWord)

	// Extract the number inside the parentheses
	insideParens := matches[1]

	// Convert the number to an integer
	number, err := strconv.Atoi(insideParens)

	if err != nil {
		return 0, err
	}
	return number, nil
}
