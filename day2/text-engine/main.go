package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main()  {
	
	if len(os.Args) != 3 {
		log.Fatal("Usage: go run . input.txt output.txt")
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if inputFile == outputFile {
		log.Fatal("✗ Both file name cannot be the same")
	}
	
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // we automatically close at the end of our program

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines) // line by line reading
	for scanner.Scan(){
		scanner.Text()
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	file2, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close() // we automatically close at the end of our program
	
	writer := bufio.NewWriter(file2)
	writer.Flush()

	fmt.Println("Program successful")

}

