package main

import(
	"fmt"
)

func main()  {
	type personality struct{
		Name 	string
		Age 	uint
		Gender 	string
		Team 	string
	}

	chibueze := new(personality)
	chibueze.Name = "Chibueze Maxwell"
	chibueze.Age = 21
	chibueze.Gender = "Male"
	chibueze.Team = "Blackspark"

	fmt.Println("Name:", chibueze.Name)
	fmt.Println("Age:", chibueze.Age)
	fmt.Println("Gender:", chibueze.Gender)
	fmt.Println("Team:", chibueze.Team)

	
}

// two methods but i prefer the new method
/*

goku := new(Saiyan)
	goku.name = "goku"
	goku.power = 9001

//vs

goku := &Saiyan {
	name: "goku",
	power: 9000,
}
*/