package main

import (
	"fmt"
)

func main() {
	strings := []string{
		"Привет",
		"Go",
		"Самая длинная строка в этом срезе",
		"Программа",
	}

	longest := ""
	for _, s := range strings {
		if len(s) > len(longest) {
			longest = s
		}
	}

	fmt.Println("Самая длинная строка:", longest)
	fmt.Scanln()
}
