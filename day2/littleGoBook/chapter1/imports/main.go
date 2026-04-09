// You basicall y need to import to go far in golang and you have to use the package you imported.

package main

import(
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("It's is over", os.Args[1])
}