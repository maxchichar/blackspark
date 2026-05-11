package render

import(
	"fmt"
	"strings"
)

func Render(asciiChar map[rune][]string, lines []string) {
	for _, line := range lines{
		if line == ""{
			fmt.Println()
			continue
		}
		
		for row := 0; row < 8; row++{
			
			var b strings.Builder

			for _, char := range line{ 
				if charRows, ok := asciiChar[char]; ok && row < len(charRows){
					b.WriteString(charRows[row])
				}
			}
			fmt.Println(b.String())
		}
	}
}