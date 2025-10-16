package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Создаём сканер для чтения строки из стандартного ввода
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите строку: ")
	
	// Считываем строку
	if scanner.Scan() {
		input := scanner.Text()
		// Преобразуем в верхний регистр
		upper := strings.ToUpper(input)
		// Выводим результат
		fmt.Println(upper)
	}
	fmt.Scanln()
	fmt.Scanln()
}
