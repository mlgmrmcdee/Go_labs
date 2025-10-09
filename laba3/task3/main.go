package main

import (
    "fmt"
    "task3/stringutils"
)

func main() {
    str := "Привет, Go!"
    reversed := stringutils.Reverse(str)
    fmt.Println("Оригинал:", str)
    fmt.Println("Перевернутая:", reversed)
}
