package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumberGenerator(ch chan int) {
	for {
		num := rand.Intn(100)
		ch <- num
		time.Sleep(1000 * time.Millisecond)
	}
}

func parityChecker(in chan int, out chan string) {
	for num := range in {
		if num%2 == 0 {
			out <- fmt.Sprintf("Число %d — чётное", num)
		} else {
			out <- fmt.Sprintf("Число %d — нечётное", num)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numCh := make(chan int)
	messageCh := make(chan string)

	go randomNumberGenerator(numCh)
	go parityChecker(numCh, messageCh)

	for {
		select {
		case msg := <-messageCh:
			fmt.Println(msg)

		case <-time.After(3 * time.Second):
			fmt.Println("Нет данных 3 секунды — завершаем работу.")
			return
		}
	}
}
