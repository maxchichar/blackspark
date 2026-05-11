package main

import(
	"fmt"
	"os"
	"strings"
)


func main() {
	usage := `
USAGE: go run . [STRING] [BANNER]
Available banners: standard, shadow thinkertoy	
	`
	if len(os.Args) < 2{
		fmt.Println("")
	}
 }