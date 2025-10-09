package main

import (
	"fmt"
)

func main() {
	// Создаем срез строк
	strings := []string{
		"Привет",
		"Go",
		"Самая длинная строка в этом срезе",
		"Программа",
	}

	// Переменные для хранения самой длинной строки
	longest := ""
	for _, s := range strings {
		if len(s) > len(longest) {
			longest = s
		}
	}

	fmt.Println("Самая длинная строка:", longest)
	fmt.Scanln()
}
