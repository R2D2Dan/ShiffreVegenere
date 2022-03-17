package solid

import (
	"fmt"
	"strings"
)

func Test(text string, key string) {
	text1 := strings.Split(text, " ")

	for _, k := range text1 {
		text2 := strings.Split(k, "")
		for i, j := range text2 {
			fmt.Println(i, j)
		}

	}
}
