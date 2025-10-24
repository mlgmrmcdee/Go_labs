package main

import (
	"fmt"
)

func CheckNumber(num int) string {
	if num > 0 {
		return "Positive"
	} else if num < 0 {
		return "Negative"
	}
	return "Zero"
}

func main() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	result := CheckNumber(number)
	fmt.Println("Результат:", result)

	fmt.Scanln()
	fmt.Scanln()
}
