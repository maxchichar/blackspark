package main

import (
	"fmt"
	"time"
)

func main()  {
	i := 1
	fmt.Println("write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// To know the day in the week
	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("It's monday")
	case time.Tuesday:
		fmt.Println("It's tuesday")
	case time.Wednesday:
		fmt.Println(time.Now(),"It's wednesday")
	case time.Thursday:
		fmt.Println("it's thursday")
	case time.Friday:
		fmt.Println("It's friday")
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	}

	// To know the hour of the day
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's afternoon")
	}

	// check the type
	WhatAmI := func (i interface{})  {
		switch t := i.(type) {
		case bool:
			fmt.Println("this is a boolean")
		case string:
			fmt.Println("this is a string value")
		default:
			fmt.Printf("The type %T is unknown\n", t)
		}
	}

	WhatAmI(true)
	WhatAmI(7)
	WhatAmI("blackspark")
}