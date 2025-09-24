// Вычисление площади прямоугольника.
package main

import "fmt"

// Структура, описывающая прямоугольник.
type Rectangle struct {
	Width  float64
	Height float64
}

// Метод возвращает площадь прямоугольника.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// Создание прямоугольника.
	rect := Rectangle{Width: 10, Height: 5}

	// Вывод площади прямоугольника.
	fmt.Println("Площадь прямоугольника:", rect.Area())

	fmt.Scanln()
}
