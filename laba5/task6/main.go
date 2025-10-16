package main

import (
	"fmt"
)

// Определяем интерфейс Stringer
type Stringer interface {
	String() string
}

// Структура Book хранит информацию о книге
type Book struct {
	Title  string
	Author string
	Year   int
}

// Реализуем интерфейс Stringer для Book
func (b Book) String() string {
	return fmt.Sprintf("Книга: \"%s\", Автор: %s, Год: %d", b.Title, b.Author, b.Year)
}

func main() {
	book := Book{
		Title:  "Преступление и наказание",
		Author: "Ф. М. Достоевский",
		Year:   1866,
	}

	var s Stringer = book
	fmt.Println(s.String())

	fmt.Scanln()
	fmt.Scanln()
}
