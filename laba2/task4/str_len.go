// Определение длины строки.
package main

import (
	"fmt"
	"unicode/utf8"
)

// Функция возвращает количество символов в строке.
func StringLength(s string) int {
	return utf8.RuneCountInString(s)
}

func main() {
	// Ввод строки с клавиатуры.
	var input string
	fmt.Print("Введите строку: ")
	fmt.Scanln(&input)

	// Вывод длины введённой строки.
	fmt.Println("Длина строки:", StringLength(input))
	
	fmt.Scanln()
	fmt.Scanln()
}
