package main

import (
	"fmt"
)

func sumAndDiff(a, b float64) (float64, float64) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func main() {
	var x, y float64

	fmt.Print("Введите первое число: ")
	fmt.Scan(&x)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&y)

	sum, diff := sumAndDiff(x, y)

	fmt.Println("Сумма:", sum)
	fmt.Println("Разность:", diff)

	fmt.Scanln()
	fmt.Scanln()
}
