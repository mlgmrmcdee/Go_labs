// Программа запрашивает два числа и выводит их сумму и разность
package main

import (
	"fmt"
)

// Функция возвращает сумму и разность двух чисел
func sumAndDiff(a, b float64) (float64, float64) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func main() {
	var x, y float64

	// Ввод двух чисел от пользователя
	fmt.Print("Введите первое число: ")
	fmt.Scan(&x)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&y)

	// Получение суммы и разности
	sum, diff := sumAndDiff(x, y)

	// Вывод результатов
	fmt.Println("Сумма:", sum)
	fmt.Println("Разность:", diff)

	fmt.Scanln()
	fmt.Scanln()
}
