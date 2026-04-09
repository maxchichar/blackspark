package main

import (
	"fmt"
)

func main()  {
	chibueze := NewChibueze("chibueze", 7)
	fmt.Printf("Name: %s Skill: %d", chibueze.Name, chibueze.Skill)
}

//usage: go run main.go constructor.go
