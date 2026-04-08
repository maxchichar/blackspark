package main

import (
	"fmt"
	"slices"
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

	a = append(a, "Lucky")
	a = append(a, "Greatman", "Ola")
	fmt.Println("apd:", a)

	c := make([]string, len(a))
	copy(c, a)
	fmt.Println("cpy:", c)

	l := a[2:5]
	fmt.Println("sl1:",l)

	l = a[:5]
	fmt.Println("sl2:",l)

	l = a[2:]
	fmt.Println("sl3:",l)
	
	t := []string{"a", "b", "c", "d"}
	fmt.Println("dlc:", t)

	t2 := []string{"a", "b", "c", "e"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}else{
		fmt.Println("t != t2")
	}

}