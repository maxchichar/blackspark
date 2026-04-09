package main

import (
	"log"
	"os"
)

func main()  {
	
	if len(os.Args) != 3 {
		log.Fatal("Usage: go run . input.txt output.txt")
		os.Exit(1)
	}

	inputFile := os.Args[2]
	outputFile := os.Args[3]

	if inputFile == outputFile {
		log.Fatal("✗ Both file name cannot be the same")
	}
	
	

}

