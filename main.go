package main

import (
	solid "ShiffreVegenere/Solid"
	"bufio"
	"fmt"
	"os"
)

func main() {

	text := bufio.NewScanner(os.Stdin)
	key := bufio.NewScanner(os.Stdin)
	language := bufio.NewScanner(os.Stdin)

	fmt.Println("        <-Шифр Веженера->")
	fmt.Printf("Введите язык в котором будет печататься.\nРусский укажите -> ru\nАнглийский -> en")
	fmt.Print("Please specify lang and [ru/en] and press enter:")
	language.Scan()
	fmt.Print("Введите текст который нужно зашифровать:")
	text.Scan()
	fmt.Print("Введите ключ который нужно зашифровать:")
	key.Scan()
	fmt.Printf("Начался процес шифрования...\n\n\n\n")

	res := solid.Solid(text.Text(), key.Text(), language.Text())

	fmt.Printf("Зашифрованный текст:%s\n", res)
	fmt.Println("Процес шифрования закончен...")

}
