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

	fmt.Println("        <-Шифр Веженера->")
	fmt.Print("Введите текст который нужно зашифровать:")
	text.Scan()
	fmt.Print("Введите ключ который нужно зашифровать:")
	key.Scan()
	fmt.Println("Начался процес шифрования...")

	res := solid.Solid(text.Text(), key.Text())

	fmt.Println(res)

}
