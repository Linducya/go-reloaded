package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if two command-line arguments are provided (input and output file names)
	if len(os.Args) != 3 {
		fmt.Println("Text modifier usage: please enter <input file> <output file>")
		os.Exit(1) // Prints exit status 1. The program did not run successfully as the required arguments were not provided.
	}

	// Variables for keywords
	// hex := "(hex)"
	// binary := "(bin)"
	// uppercase := "(up)"
	// lowercase := "(low)"
	// caps := "(cap)"
	// capNumber := 0

	// Get the input & output file names from command-line arguemts
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	fmt.Println(inputFile)
	fmt.Println(outputFile)

	// Read the contents from the input file into a []byte var inputData
	inputData, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err) // Error reading input file: open sample.txt: no such file or directory
		os.Exit(1)
	}

	// fmt.Println(inputData) // Prints the data in the array of []byte taken from the input file
	fmt.Println("Original Content:")
	fmt.Println(string(inputData)) // Print the string of inputData

	// Modify the contents

	// Convert to uppercase with for loop
	// for i := 0; i < len(inputData); i++ {
	// 	if inputData[i] >= 'a' && inputData[i] <= 'z' {
	// 		inputData[i] -= 'a' - 'A'
	// 	}
	// }
	// modifiedData := inputData

	// Convert to uppercase with strings.ToUpper():
	modifiedData := strings.ToUpper(string(inputData))
	fmt.Println("Modified content")
	fmt.Println(string(modifiedData))

	// // Convert the string to slice of byte
	// outputData := []byte(modifiedData)
	// fmt.Println(string(outputData))

	// Write the modified contents to the output file, 0644 6=rw Owner, 4=r Group, Others
	err = os.WriteFile(outputFile, []byte(modifiedData), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(1)
	}

}
