package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Agent: Hello world!")

	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path as an argument.")
		os.Exit(1)
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		os.Exit(1)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		os.Exit(1)
	}

	fileContent := string(data)

	fmt.Println(fileContent)
}
