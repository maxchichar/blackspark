// Based on my understanding i believe that a slice is a flexible array and it enables you to manipulate.

package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// (low)
func ToLowerLastWord(word []string) []string{
	if len(word) == 0 {
		return word
	}
	if word[len(word) - 1] == "(low)" {
		word[len(word) - 2] = strings.ToLower(word[len(word)-2])
		word = word[:len(word)-1]
	}

	return word
}

// (up)
func ToUpperLastWord(text []string) []string {
	if len(text) == 0 {
		return text
	}

	if text[len(text)-1] == "(up)" {
		text[len(text)-2] = strings.ToUpper(text[len(text)-2])
		text = text[:len(text)-1]
	}
	return text
}

//(cap)
func ToCapLastWord(text []string) []string {
	if len(text) == 0 {
		return text
	}

	if text[len(text)-1] == "(cap)" {
		text[len(text)-2] = cases.Title(language.English).String(text[len(text)-2])
		text = text[:len(text)-1]
	}
	return text
}

func LastIndex(word []string, index int) []string {
	lastIndex := len(word)-1

	word[index], word[lastIndex] = word[lastIndex], word[index]
	word = word[:len(word)-1]
	return word
}

func to_upper(text []string, n int) []string {
	lastIndex := len(text) - 1 

	for i := 0; i < lastIndex; i++ {
		if text[i] == "(up)" {
			text[len(text)-n] = cases.Upper(language.English).String(text[len(text)-n])
			text[i], text[lastIndex] = text[lastIndex], text[i]
			text = text[:lastIndex]
			// if text[i] == "(up)" {
			// 	text[i] = strings.ReplaceAll(text[i], "(up)", "")
			// }
		}
	}
	return text
}

func main()  {
	scores := []int{1,2,3,4,5,6}
	fmt.Println(scores)

	scores1 := make([]int, 10) // we also use make in creating a slice
	fmt.Println(scores1)
/*
	score2 := make([]int, 0, 100) // has the capacity of hundred with a length of zero
	score2[77] = 100
	fmt.Println(score2) // this will print a runtime error because the lenght is zero 
*/
	// the only way to solve this is to use append because it expands our slice
	score3 := make([]int, 0, 10)
	score3 = append(score3, 5)
	fmt.Println(score3)

	fmt.Println(ToLowerLastWord([]string{"GO", "IS", "FUN", "(low)"}))
	fmt.Println(ToUpperLastWord([]string{"hello", "world", "(up)"}))

	result := []string{"master", "sparkyechox", "(cap)"}
	fmt.Println(ToCapLastWord(result))

	word := []string{"my", "name", "is", "Chibueze"}
	output := LastIndex(word, 3) // removes the index you want just by changing the index number
	fmt.Println(output)

	text := []string{"i", "love", "(up)", "golang"}
	out := to_upper(text, 3)
	fmt.Println(out)
}