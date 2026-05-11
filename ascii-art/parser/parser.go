package parser

import(
	"strings"
)

func Parser(data string) []string{
	if data == "\\n"{
		clean := []string{" "}
		return clean
	}

	if data == ""{
		empty := []string{""}
		return empty
	}

	return strings.Split(strings.ReplaceAll(data, "\\n", "\n"), "\n")
}