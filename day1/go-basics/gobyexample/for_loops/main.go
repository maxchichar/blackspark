package main

import(
	"fmt"
)

func main()  {
	i := 1
	for i <= 3{
		fmt.Println(i)
		i = i + 1
	}

	// for calculating integers
	for j := 0; j < 3; j++ {
		fmt.Println("loops:", j)
	}

	for i := range 3{
		fmt.Println("range:", i)
	}

	// for calculating string
	word := "hello"
	for v := range word {
		fmt.Println(v)
	}

	for l := 0; l < len(word); l++{
		fmt.Println(l)
	}

	// breaking loops
	for {
		fmt.Println("loops")
		break
	}

	// divide to get the odd numbers
	for n := range 10{
		if n%2 == 0{
			fmt.Println(n)
		}
		continue
	}
}