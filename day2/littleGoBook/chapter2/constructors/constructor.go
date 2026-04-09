package main

type chibueze struct{
	Name string
	Skill int
}


func NewChibueze(name string, skill int) chibueze {
	return chibueze{
		Name: name,
		Skill: skill,
	}
}

// func main()  {
// 	chibueze := NewChibueze("chibueze", 7)
// 	fmt.Printf("Name: %s Skill: %d", chibueze.Name, chibueze.Skill)
// }