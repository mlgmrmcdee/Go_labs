package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"laba3/task1/mathutils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите число: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	n, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	res, err := mathutils.Factorial(n)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("Факториал %d = %s\n", n, res.String())
}
