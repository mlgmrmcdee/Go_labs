package main

import (
	"fmt"
)

func main() {
	var a, b, c float64

	fmt.Print("Введите три числа через пробел: ")
	fmt.Scan(&a, &b, &c)

	average := (a + b + c) / 3

	fmt.Printf("Среднее значение: %.2f\n", average)

	fmt.Scanln()
	fmt.Scanln()
}
