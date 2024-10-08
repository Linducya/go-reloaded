package textutils

import (
	"fmt"
	"regexp"
)

func FindPatterns(word string) (string, error) {
	// Compile a regex pattern to capture the keyword (cap, up, low)
	pattern := `\((cap|up|low),`

	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Find the match and capture the number
	matches := re.FindStringSubmatch(word)

	// Check if the match was found
	if len(matches) > 1 {
		// Extract the keyword(cap, up, low)
		keyword := matches[1]
		return keyword, nil
	}

	// Return an error if no match found
	return "", fmt.Errorf("no keyword found")
}
