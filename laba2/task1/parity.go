package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число нечётное")
	}

	fmt.Scanln()
	fmt.Scanln()
}
