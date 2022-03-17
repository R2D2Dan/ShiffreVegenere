package solid

import (
	"strings"
)

func Solid(text string, key string) string {
	alphabetslice := strings.Split("абвгдеёжзийклмнопрстуфхцчшщъыьэюя", "")
	keyslice := strings.Split(key, "")
	textslice := strings.Split(text, " ")
	countspace := len(textslice) - 1
	res := ""

	for _, k := range textslice {
		text2 := strings.Split(k, "")
		for i, j := range text2 {

			textindex := SearchIndex(alphabetslice, string(j))
			keyindex := SearchIndex(alphabetslice, string(keyslice[i%len(keyslice)]))
			res = res + alphabetslice[(textindex+keyindex)%len(alphabetslice)]
		}
		if countspace > 0 {
			res = res + " "
			countspace--
		}

	}
	return res

}

// Поиск индекса в строке
func SearchIndex(alphabet []string, char string) int {
	for i, k := range alphabet {
		if k == char {
			return i
		}
	}
	return -1
}
