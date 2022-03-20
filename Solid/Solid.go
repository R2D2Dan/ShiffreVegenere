package solid

import (
	"strings"
)

func Solid(text string, key string, language string) []string {

	ru := "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	en := "abcdefghijklmnopqrstuvwxyz"

	var alphabetslice []string
	switch language {
	case "ru":
		alphabetslice = strings.Split(ru, "")
	case "en":
		alphabetslice = strings.Split(en, "")
	}

	var res []string
	textclice := strings.Split(text, "")
	keyclice := strings.Split(key, "")
	i := 0
	for _, k := range textclice {
		if string(k) != " " {

			indextext := SearchIndex(alphabetslice, string(k))

			keyindex := SearchIndex(alphabetslice, keyclice[i%len(keyclice)])

			res = append(res, alphabetslice[(indextext+keyindex)%len(alphabetslice)])
			i++
		} else {
			res = append(res, " ")
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
