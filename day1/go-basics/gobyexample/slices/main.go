package main

import(
	"fmt"
)

func main()  {
	
	// Any slice without a value/uninitialised is equals to zero
	var a []string
	fmt.Println("uninit:", a, a == nil, len(a) == 0)

	// To make a slice that has a non-zero length we use make
	a = make([]string, 3)
	fmt.Println("emp:", a, "len:", len(a), "cap:", cap(a))

	a[0] = "Chibueze"
	a[1] = "Charles"
	a[2] = "Maxwell"
	fmt.Println("set:", a)
	fmt.Println("get:", a[1])

	fmt.Println(len(a))
}