package main

import (
	"fmt"
	"unicode/utf8"
)

func StringLength(s string) int {
	return utf8.RuneCountInString(s)
}

func main() {
	var input string
	fmt.Print("Введите строку: ")
	fmt.Scanln(&input)

	fmt.Println("Длина строки:", StringLength(input))

	fmt.Scanln()
	fmt.Scanln()
}
