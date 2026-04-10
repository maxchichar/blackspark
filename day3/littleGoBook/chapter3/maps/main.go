// Maps in Go are like what others call hashtables or dictionaries.
// Maps are also created / declared using make function
package main

import (
	"fmt"
	"strings"
)

func main()  {
	text := "chibueze"

	lookup := make(map[string]string)
	lookup ["(up)"] = strings.ToUpper(text)
	
	skill, exists := lookup["(up)"]

	fmt.Println(skill, exists)

	total := len(lookup) // check the length of the keys

	fmt.Println(total)

	//We need to set an initial size to know the size of our key
	num := 21
	check := make(map[string]int, 100)
	check["king"] = num
	tac, exist := check["queen"]
	fmt.Println(tac, exist) 

	// declaring maps as a struct
	type chibueze struct{
		Name string
		Skill map[string]*chibueze
	}

	sparkyechox := &chibueze{
		Name: "Prince",
		Skill: make(map[string]*chibueze),
	}

	fmt.Println(sparkyechox.Name)
	sparkyechox.Skill["Golang"] = &chibueze{
		Name: "Golang programming",
		Skill: make(map[string]*chibueze),
	}

	fmt.Println(len(sparkyechox.Skill))
}