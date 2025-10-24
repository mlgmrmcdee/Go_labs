package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func worker(id int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for line := range jobs {
		reversed := reverseString(line)
		results <- reversed
	}
}

func main() {
	start := time.Now()

	inputFile := "input.txt"
	outputFile := "output.txt"

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	var workerCount int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&workerCount)

	jobs := make(chan string, len(lines))
	results := make(chan string, len(lines))

	var wg sync.WaitGroup

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for _, line := range lines {
		jobs <- line
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Ошибка при создании выходного файла:", err)
		return
	}
	defer out.Close()

	for res := range results {
		out.WriteString(res + "\n")
	}

	duration := time.Since(start)
	fmt.Printf("Результаты записаны в %s\n", outputFile)
	fmt.Printf("Время выполнения: %v\n", duration)
}
