package main

import(
	"fmt"
)

func main() {
	type chibueze struct{
		Id int
		Pfs string
	}

	name := chibueze{
		Id: 19,
		Pfs: "student",
	}
	

	fmt.Println(name.Id)
	fmt.Println(name.Pfs)
}