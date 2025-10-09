package main

import (
	"fmt"
)

func main() {
	// Создадим массив и возьмём срез
	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:4] // элементы с индекса 1 по 3 (20, 30, 40)
	fmt.Println("Исходный срез:", slice)

	// Добавление элемента
	slice = append(slice, 60)
	fmt.Println("После добавления:", slice)

	// Удаление элемента
	index := 1
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println("После удаления:", slice)
}
