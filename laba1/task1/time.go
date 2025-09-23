package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println("Текущая дата и время:", now.Format("2006-01-02 15:04:05"))
}
