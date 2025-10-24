package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		return
	}
	fmt.Println("Сервер запущен и слушает порт 8080...")

	var wg sync.WaitGroup
	connections := make(map[net.Conn]bool)
	mu := sync.Mutex{}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	shutdown := make(chan struct{})
	go func() {
		<-stop
		fmt.Println("\nОстановка сервера...")
		listener.Close() // Закроет Accept и вызовет ошибку
		close(shutdown)  // Уведомим главный цикл
		mu.Lock()
		for conn := range connections {
			conn.Close()
		}
		mu.Unlock()
		fmt.Println("Все соединения закрыты. Ожидание завершения горутин...")
	}()
acceptLoop:
	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-shutdown:
				break acceptLoop
			default:
				fmt.Println("Ошибка при подключении клиента:", err)
				continue
			}
		}

		mu.Lock()
		connections[conn] = true
		mu.Unlock()

		wg.Add(1)
		go handleConnection(conn, &wg, &mu, connections)
	}

	wg.Wait()
	fmt.Println("Сервер корректно завершил работу.")
}

func handleConnection(conn net.Conn, wg *sync.WaitGroup, mu *sync.Mutex, connections map[net.Conn]bool) {
	defer wg.Done()
	defer func() {
		mu.Lock()
		delete(connections, conn)
		mu.Unlock()
		conn.Close()
	}()

	fmt.Println("Клиент подключен:", conn.RemoteAddr())
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Клиент отключился:", conn.RemoteAddr())
			return
		}

		fmt.Printf("[%s]: %s", conn.RemoteAddr(), message)

		_, err = conn.Write([]byte("Сообщение получено!\n"))
		if err != nil {
			fmt.Println("Ошибка при отправке ответа:", err)
			return
		}
	}
}
