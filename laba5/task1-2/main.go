package main

import "fmt"

// Определяем структуру Person
type Person struct {
	Name string
	Age  int
}

// Метод для вывода информации о человеке
func (p Person) Info() {
	fmt.Printf("Имя: %s, Возраст: %d\n", p.Name, p.Age)
}

// Метод, увеличивающий возраст на 1 год
func (p *Person) Birthday() {
	p.Age++
}

func main() {
	person := Person{Name: "Алексей", Age: 25}

	person.Info() // Выводим начальную информацию

	person.Birthday() // Отмечаем день рождения

	person.Info() // Проверяем, что возраст увеличился

	fmt.Scanln()
	fmt.Scanln()
}
