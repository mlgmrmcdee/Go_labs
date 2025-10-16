package main

import (
	"fmt"
)

func main() {
	var num, sum int

	fmt.Println("Введите числа (0 для завершения):")

	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		sum += num
	}

	fmt.Println("Сумма введённых чисел:", sum)
	fmt.Scanln()
	fmt.Scanln()
}
