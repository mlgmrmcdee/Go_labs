package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите строку: ")

	if scanner.Scan() {
		input := scanner.Text()
		upper := strings.ToUpper(input)
		fmt.Println(upper)
	}
	fmt.Scanln()
	fmt.Scanln()
}
