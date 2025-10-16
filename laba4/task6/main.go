package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println("Исходный массив:", arr)

	// Переворачиваем массив
	reversed := make([]int, len(arr))
	for i := range arr {
		reversed[i] = arr[len(arr)-1-i]
	}

	fmt.Println("Массив в обратном порядке:", reversed)
	fmt.Scanln()
	fmt.Scanln()
}
