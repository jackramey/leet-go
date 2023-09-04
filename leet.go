package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if an argument is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run script.go INPUT_STRING")
		return
	}

	input := os.Args[1]

	// Split the input string using the underscore character
	parts := strings.Split(input, "_")

	// Extract the problem number and description
	problemNumber := parts[0]
	problemDesc := parts[1]

	// Pad the problem number with leading zeros to make it four digits long
	paddedProblemNumber := fmt.Sprintf("%04s", problemNumber)

	// Create the problem name
	problemName := paddedProblemNumber + "_" + problemDesc

	// Create the directory
	err := os.MkdirAll(problemName, os.ModePerm)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create the Go file and write the package statement
	goFile := problemName + "/" + problemName + ".go"
	file, err := os.Create(goFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "package %s\n", problemDesc)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
