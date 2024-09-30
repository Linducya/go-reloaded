package textutils

import "strconv"

// Helper function to convert hex to decimal
func hexToDecimal(hex string) (int64, error) {
	// Convert the hex strting to decimal (base 10)
	decimal, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return 0, err
	}
	return decimal, nil
}
