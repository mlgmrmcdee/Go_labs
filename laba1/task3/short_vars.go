// Пример короткого объявления переменных и вывода их значений
package main

import "fmt"

func main() {
    // Короткое объявление переменных с автоматическим определением типа
    age := 20
    pi := 3.14159
    name := "mlgmrmcdee"
    isGoCool := true

    // Вывод значений в консоль
    fmt.Println("Целое число (int):", age)
    fmt.Println("Число с плавающей точкой (float64):", pi)
    fmt.Println("Строка (string):", name)
    fmt.Println("Булево значение (bool):", isGoCool)

    fmt.Scanln()
}
