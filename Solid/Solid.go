package solid

import (
	"fmt"
	"strings"
)

func Solid(text string, key string, language string) []string {

	ru := "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	en := "abcdefghijklmnopqrstuvwxyz"
	var alphabetslice []string
	var res []string

	switch language {
	case "ru":
		alphabetslice = strings.Split(ru, "")
	case "en":
		alphabetslice = strings.Split(en, "")
	default:
		return []string{"Ошибка"}
	}

	keyslice := strings.Split(key, "")
	textslice := strings.Split(text, " ")
	fmt.Println(textslice)
	for _, k := range textslice {
		text2 := strings.Split(k, "")
		word := ""
		for i, j := range text2 {

			textindex := SearchIndex(alphabetslice, string(j))
			keyindex := SearchIndex(alphabetslice, string(keyslice[i%len(keyslice)]))
			word = word + alphabetslice[(textindex+keyindex)%len(alphabetslice)]

		}
		fmt.Println(word)

		res = append(res, word)

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
