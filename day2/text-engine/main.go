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

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if inputFile == outputFile {
		log.Fatal("✗ Both file name cannot be the same")
	}
	
	file, err := os.OpenFile("sample.txt", os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal(err)
	}
	if err := file.Close(); err != nil{ // we automatically close at the end of our program
		log.Fatal(err)
	}
	
	
	file2, err := os.OpenFile("result.txt", os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal(err)
	}
	if err := file2.Close(); err != nil{ // we automatically close at the end of our program
		log.Fatal(err)
	}

}

