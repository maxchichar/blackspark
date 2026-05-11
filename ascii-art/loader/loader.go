package loader

import(
	"os"
	"io"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	fileOpen, err := os.Open(filename)
	if err != nil {
		return err, nil
	}
	defer fileOpen.Close()

	fileRead, err := io.ReadAll(fileOpen)
	if err != nil{
		return err, nil
	}

	fileStr := strings.ReplaceAll(string(fileRead), "\r\n", "\n")
	file := strings.Split(strings.Trim(fileStr, "\n"), "\n\n")

	charMap := make(map[rune][]string)

	for i, fileChar := range file{
		r := rune(32 + i)
		charMap[r] = strings.Split(strings.Trim(fileChar, "\n"), "\n")
	}
	return charMap, nil
}