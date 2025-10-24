package main

import (
	"fmt"
)

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	fmt.Println("Исходный срез:", slice)

	slice = append(slice, 60)
	fmt.Println("После добавления:", slice)

	index := 1
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println("После удаления:", slice)

	fmt.Scanln()
}
