package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	fmt.Printf("\nРезультаты:\n")
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)

	if b != 0 {
		fmt.Printf("%d / %d = %d\n", a, b, a/b)
		fmt.Printf("%d %% %d = %d\n", a, b, a%b)
	} else {
		fmt.Println("Деление и остаток невозможны (деление на ноль).")
	}
}
