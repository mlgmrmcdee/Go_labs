package main

import "fmt"

func Average(a, b int) float64 {
	return float64(a+b) / 2.0
}

func main() {
	var x, y int
	fmt.Print("Введите два числа через пробел: ")
	fmt.Scan(&x, &y)

	fmt.Println("Среднее значение:", Average(x, y))

	fmt.Scanln()
	fmt.Scanln()
}
