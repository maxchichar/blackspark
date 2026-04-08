package main

import(
	"fmt"
)

func main()  {
	var a [5]int
	fmt.Println("emp:", a)
	a[3] = 77
	a[4] = 100
	fmt.Println("set", a)
	fmt.Println("get", a[4])
	fmt.Println(len(a))
	
	
	var b [6]string
	b = [6]string{"porsche 911", "lamborghini", "ferrari", "toyota", "xiaomi", "G-wagon"}
	fmt.Println(b)
	fmt.Println(len(b))
	
	//getting the word in the array
	g := b[1]
	fmt.Println("get:", g)
	fmt.Println(len(g))

	c := [7]int{1,2,3,4,5,6,7}
	fmt.Println(c)

	c = [...]int{1,2,3,4,5,6,7}
	fmt.Println(c)

	c = [...]int{100, 5: 400, 500}
	fmt.Println(c)

	var twoD [2][4]int
	for i := range 2{
		for r := range 4{
			twoD[i][r] = i + r //adding both arrays
		}
	}
	fmt.Println("2d", twoD)

	twoD = [2][4]int{ // swaping
		{1,2,3,4},
		{1,2,3,5},
	}
	fmt.Println("2d", twoD)
}
