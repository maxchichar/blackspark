package main

import(
	"fmt"
)

func getPower() int {
	return 7000
}

func main() {
	power := getPower()
	fmt.Println("power =", power)
}