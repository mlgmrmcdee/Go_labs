package main

import (
	"fmt"
	"math"
)

// Интерфейс Shape с методом Area()
type Shape interface {
	Area() float64
}

// Структура прямоугольника
type Rectangle struct {
	Width, Height float64
}

// Реализация метода Area() для Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Структура круга
type Circle struct {
	Radius float64
}

// Реализация метода Area() для Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Функция, принимающая срез Shape и выводящая площади всех фигур
func PrintAreas(shapes []Shape) {
	for i, shape := range shapes {
		fmt.Printf("Фигура %d: площадь = %.2f\n", i+1, shape.Area())
	}
}

func main() {
	// Создание фигур
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 3}

	// Срез интерфейсов Shape
	shapes := []Shape{rect, circle}

	// Вызов функции для вывода площадей
	PrintAreas(shapes)

	fmt.Scanln()
	
}
