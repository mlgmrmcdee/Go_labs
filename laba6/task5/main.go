package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Request struct {
	A, B   float64
	Op     rune
	RespCh chan Result
	Id     int
}

type Result struct {
	Value float64
	Err   error
	Id    int
}

func worker(id int, reqCh <-chan Request, wg *sync.WaitGroup) {
	defer wg.Done()
	for req := range reqCh {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		var res Result
		res.Id = req.Id

		switch req.Op {
		case '+':
			res.Value = req.A + req.B
		case '-':
			res.Value = req.A - req.B
		case '*':
			res.Value = req.A * req.B
		case '/':
			if req.B == 0 {
				res.Err = errors.New("деление на ноль")
			} else {
				res.Value = req.A / req.B
			}
		default:
			res.Err = fmt.Errorf("неизвестная операция: %q", req.Op)
		}

		req.RespCh <- res
	}
}

func sendRequest(reqCh chan<- Request, a, b float64, op rune, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	resp := make(chan Result)
	req := Request{
		A:      a,
		B:      b,
		Op:     op,
		RespCh: resp,
		Id:     id,
	}

	reqCh <- req

	result := <-resp

	if result.Err != nil {
		fmt.Printf("Запрос %2d: %.2f %c %.2f → ошибка: %v\n", id, a, op, b, result.Err)
	} else {
		fmt.Printf("Запрос %2d: %.2f %c %.2f = %.2f\n", id, a, op, b, result.Value)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	requests := make(chan Request)

	numWorkers := 4
	var workerWg sync.WaitGroup
	workerWg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i, requests, &workerWg)
	}

	var clientsWg sync.WaitGroup
	numClients := 12
	ops := []rune{'+', '-', '*', '/'}
	clientsWg.Add(numClients)

	for i := 0; i < numClients; i++ {
		go func(id int) {
			defer clientsWg.Done()
			a := float64(rand.Intn(20) - 5)
			b := float64(rand.Intn(11) - 5)
			op := ops[rand.Intn(len(ops))]

			var innerWg sync.WaitGroup
			innerWg.Add(1)
			go sendRequest(requests, a, b, op, id, &innerWg)
			innerWg.Wait()
		}(i + 1)
	}

	clientsWg.Wait()

	close(requests)

	workerWg.Wait()

	fmt.Println("Работа калькулятора завершена.")
}
