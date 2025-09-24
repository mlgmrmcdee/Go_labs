// Вычисление среднего арифметического двух чисел.
package main

import "fmt"

// Функция возвращает среднее арифметическое двух целых чисел.
func Average(a, b int) float64 {
	return float64(a+b) / 2.0
}

func main() {
	// Ввод двух чисел с клавиатуры.
	var x, y int
	fmt.Print("Введите два числа через пробел: ")
	fmt.Scan(&x, &y)

	// Вывод среднего значения.
	fmt.Println("Среднее значение:", Average(x, y))

	fmt.Scanln()
	fmt.Scanln()
}
