package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: lat \"text\"")
		return
	}

	fmt.Println(ToLat(os.Args[1]))
}

func ToLat(text string) string {
	m := map[string]string{
		"г": "g", "д": "d", "и": "i", "й": "j", "м": "m", "ц": "c", "ы": "y", "к": "k", "ю": "ju", "е": "e", "л": "l", "ф": "f", "х": "x", "ь": "ʹ", "ж": "ž", "р": "r", "у": "u", "э": "è", "в": "v", "с": "s", "щ": "šč", "ъ": "ʺ", "ё": "ë", "о": "o", "п": "p", "я": "ja", "н": "n", "т": "t", "ч": "č", "а": "a", "б": "b", "з": "z", "ш": "š", "А": "A", "Б": "B", "В": "V", "Г": "G", "Д": "D", "Е": "E", "Ё": "Ë", "Ж": "Ž", "З": "Z", "И": "I", "Й": "J", "К": "K", "Л": "L", "М": "M", "Н": "N", "О": "O", "П": "P", "Р": "R", "С": "S", "Т": "T", "У": "U", "Ф": "F", "Х": "X", "Ц": "C", "Ч": "Č", "Ш": "Š", "Щ": "Šč", "Ъ": "ʺ", "Ы": "Y", "Ь": "ʹ", "Э": "È", "Ю": "Ju", "Я": "Ja",
	}

	output := ""

	for _, w := range strings.Split(text, "") {
		if l, ok := m[w]; ok {
			output += l
		} else {
			output += w
		}
	}

	return output
}
