package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Info() {
	fmt.Printf("Имя: %s, Возраст: %d\n", p.Name, p.Age)
}

func (p *Person) Birthday() {
	p.Age++
}

func main() {
	person := Person{Name: "Алексей", Age: 25}
	person.Info()

	person.Birthday()

	person.Info()

	fmt.Scanln()
	fmt.Scanln()
}
