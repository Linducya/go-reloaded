package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Check if two command-line arguments are provided (input and output file names)
	if len(os.Args) != 3 {
		fmt.Println("Text modifier usage: please enter <input file> <output file>")
		os.Exit(1) // Exit status 1, program did not run successfully - the required arguments were not provided.
	}

	// Get the input & output file names from command-line arguemts
	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	fmt.Println(inputFileName)
	fmt.Println(outputFileName)

	// Read the contents of the input file
	inputData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	fmt.Println("Original Content:")
	fmt.Println(string(inputData))

}
