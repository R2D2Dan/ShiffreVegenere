package solid

import (
	"strings"
)

func Solid(text string, key string) string {
	alphabet := "абвгдеёжзийклмнопрстуфхцчшщъыьэюя "
	alphabetrune := []rune(alphabet)
	keyrune := []rune(key)
	res := ""
	for i, k := range text {
		textindex := strings.Index(alphabet, string(k))/2 + (strings.Index(alphabet, string(k)) % 2)
		keyindex := strings.Index(alphabet, string(keyrune[(i/2)%len(keyrune)])) / 2
		res = res + string(alphabetrune[(textindex+keyindex)%len(alphabetrune)])
	}

	return res

}
