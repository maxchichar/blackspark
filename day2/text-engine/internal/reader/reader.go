package reader

import (
	"log"
	"os"
)

func Reader(name string) ([]byte, error) {
	file, err := os.ReadFile("text-engine/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(file)
}