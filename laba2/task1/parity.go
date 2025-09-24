// Определение чётности числа.
package main

import (
	"fmt"
)

func main() {
	// Ввод числа с клавиатуры.
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	// Проверка числа на чётность.
	if number%2 == 0 {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число нечётное")
	}

	fmt.Scanln()
	fmt.Scanln()
}
