// Демонстрация объявления переменных разных типов и их вывод
package main

import "fmt"

func main() {
    // Объявление и инициализация переменных
    var age int = 20            // целое число
    var pi float64 = 3.14159    // число с плавающей точкой
    var name string = "mlgmrmcdee" // строка
    var isGoCool bool = true    // булево значение

    // Вывод значений в консоль
    fmt.Println("Целое число (int):", age)
    fmt.Println("Число с плавающей точкой (float64):", pi)
    fmt.Println("Строка (string):", name)
    fmt.Println("Булево значение (bool):", isGoCool)

    fmt.Scanln()
}
