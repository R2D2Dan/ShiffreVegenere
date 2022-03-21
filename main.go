package main

import (
	solid "ShiffreVegenere/Solid"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("        <-Шифр Веженера->")
	fmt.Printf("Введите язык в котором будет печататься.\nРусский\t   -> ru\nАнглийский -> en\n")
	fmt.Print("Please specify lang and [ru/en] and press enter:")

	solidtext, unsolidtext := Start()
	fmt.Printf("\nЗашифрованный текст:%v ", solidtext)
	fmt.Printf("\nРасшиврованный текст:%v ", unsolidtext)

	fmt.Println("\nПроцес шифрования закончен...")

}

func Start() ([]string, []string) {

	words := bufio.NewScanner(os.Stdin)
	key := bufio.NewScanner(os.Stdin)
	language := bufio.NewScanner(os.Stdin)

	language.Scan()
	for {
		if language.Text() == "en" || language.Text() == "ru" {
			break
		}

		fmt.Println("Ошибка ввода языка попробуйте еще раз")
		fmt.Print("Please specify lang and [ru/en] and press enter:")
		language.Scan()

	}

	fmt.Printf("\nВведите текст который нужно зашифровать:")
	words.Scan()
	fmt.Printf("\nВведите ключ который нужно зашифровать:")
	key.Scan()
	fmt.Printf("Начался процес шифрования...\n\n")

	solidtext := solid.Solid(words.Text(), key.Text(), language.Text())
	unsolidtext := solid.Unsolid(strings.Join(solidtext, ""), key.Text(), language.Text())

	return solidtext, unsolidtext

}
