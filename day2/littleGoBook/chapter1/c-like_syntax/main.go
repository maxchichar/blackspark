package main

import "fmt"

func main()  {
	var words string
	words = "(up)"
	name := "Sung Jiwoo"
	if words == "(up)" {
		println("letter capitalised")
	}

	var power int
	fmt.Println("What is Sung jiwoo power: ")
	fmt.Scan(&power)
	if (name == "Sung Jiwoo" && power > 9000) || (name == "Goku" && power < 3000) {
		fmt.Println("system")
	}

}