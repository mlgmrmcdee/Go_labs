package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// инициализируем генератор случайных чисел текущим временем
	rand.Seed(time.Now().UnixNano())

	// создаем массив из 5 целых чисел
	var arr [5]int

	// заполняем массив случайными числами от 0 до 99
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) // число в диапазоне [0, 99]
	}

	// выводим массив на экран
	fmt.Println("Массив случайных чисел:")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
