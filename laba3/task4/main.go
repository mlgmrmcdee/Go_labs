package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var arr [5]int

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}

	fmt.Println("Массив случайных чисел:")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Scanln()
}
