package main

import "fmt"

// Функция для вычисления среднего возраста
func averageAge(people map[string]int) float64 {
	if len(people) == 0 {
		return 0
	}

	var sum int
	for _, age := range people {
		sum += age
	}
	return float64(sum) / float64(len(people))
}

func main() {
	// Создаем карту с именами и возрастами
	people := map[string]int{
		"Алексей": 25,
		"Мария":   30,
		"Иван":    22,
	}

	// Выводим исходный список
	fmt.Println("Исходный список людей:")
	for name, age := range people {
		fmt.Printf("%s: %d лет\n", name, age)
	}

	// Добавляем нового человека
	var newName string
	var newAge int
	fmt.Print("\nВведите имя нового человека: ")
	fmt.Scan(&newName)
	fmt.Print("Введите возраст: ")
	fmt.Scan(&newAge)

	people[newName] = newAge

	// Выводим обновлённый список
	fmt.Println("\nСписок после добавления нового человека:")
	for name, age := range people {
		fmt.Printf("%s: %d лет\n", name, age)
	}

	// Удаляем запись по имени
	var nameToDelete string
	fmt.Print("\nВведите имя для удаления: ")
	fmt.Scan(&nameToDelete)

	if _, exists := people[nameToDelete]; exists {
		delete(people, nameToDelete)
		fmt.Printf("Запись '%s' удалена.\n", nameToDelete)
	} else {
		fmt.Printf("Человек с именем '%s' не найден.\n", nameToDelete)
	}

	// Выводим итоговый список
	fmt.Println("\nИтоговый список людей:")
	for name, age := range people {
		fmt.Printf("%s: %d лет\n", name, age)
	}

	// Вычисляем и выводим средний возраст
	avg := averageAge(people)
	fmt.Printf("\nСредний возраст: %.2f лет\n", avg)
	fmt.Scanln()
	fmt.Scanln()
	fmt.Scanln()
}
