// dynamic arrays are arrays that resize based on the data that is added to them. but in Go arrays are fixed and cannot be flexible.

package main

import(
	"fmt"
)

func main()  {
	var scores [10]int
	scores[0] = 339

	fmt.Println(scores[0])

	//initializing arrays with value
	marks := [5]int{1,2,3,4,5}
	fmt.Println(marks)
	// you can loop through the arrays to find the length
	for value := range marks{
		fmt.Println("Range:",value)
	}
	
	words := [4]string{"cars", "mansion", "assets", "cities"}
	for i := 0; i < len(words); i++{
		fmt.Println("Loops:", i)
	}
}