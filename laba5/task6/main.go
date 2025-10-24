package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Book struct {
	Title  string
	Author string
	Year   int
}

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
