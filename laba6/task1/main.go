package main

import (
	"fmt"
	"math/rand"
	"time"
)

func factorial(n int) {
	fmt.Printf("Начало вычисления факториала(%d)\n", n)
	time.Sleep(2 * time.Second)
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Printf("Факториал(%d) = %d\n", n, result)
}

func randomNumbers(count int) {
	fmt.Printf("Начало генерации %d случайных чисел\n", count)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		num := rand.Intn(100)
		fmt.Printf("Случайное число %d: %d\n", i+1, num)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Генерация случайных чисел завершена")
}

func sumSeries(n int) {
	fmt.Printf("Начало вычисления суммы ряда до %d\n", n)
	time.Sleep(1 * time.Second)
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("Сумма ряда до %d = %d\n", n, sum)
}

func main() {
	fmt.Println("Запуск горутин...")

	go factorial(5)
	go randomNumbers(5)
	go sumSeries(10)

	time.Sleep(5 * time.Second)
	fmt.Println("Главная функция завершена")
}
