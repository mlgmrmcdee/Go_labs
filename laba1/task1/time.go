// Вывод текущей даты и времени в формате "YYYY-MM-DD HH:MM:SS"
package main

import (
	"fmt"
	"time"
)

func main() {
	// Получаем объект с текущими датой и временем.
	now := time.Now()

	// Выводим строку с текущей датой и временем.
	fmt.Println("Текущая дата и время:", now.Format("2006-01-02 15:04:05"))
	
	fmt.Scanln()
}
