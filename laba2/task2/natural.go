// Определение знака числа: положительное, отрицательное или ноль.
package main

import (
	"fmt"
)

// Функция возвращает строку в зависимости от значения числа.
func CheckNumber(num int) string {
	if num > 0 {
		return "Positive"
	} else if num < 0 {
		return "Negative"
	}
	return "Zero"
}

func main() {
	// Ввод числа с клавиатуры.
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	// Вызов функции и вывод результата.
	result := CheckNumber(number)
	fmt.Println("Результат:", result)

	fmt.Scanln()
	fmt.Scanln()
}
