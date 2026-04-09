package main

import(
	"fmt"
)

func main() {
	type maxchichar struct{
		Name 	string
		Skills 	int
		Father 	*maxchichar
	}

	sparkyechox := maxchichar{
		Name: "Chibueze",
		Skills: 7,
		Father: &maxchichar{
			Name: "Maxwell",
			Skills: 21,
			Father: nil,
		},
	}

	fmt.Printf("My name is %s and i have over %d skills while my dad %s has over %d skills", sparkyechox.Name, sparkyechox.Skills, sparkyechox.Father.Name, sparkyechox.Father.Skills)
}