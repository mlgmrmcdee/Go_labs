package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка при подключении к серверу:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Подключено к серверу. Введите сообщение (или 'exit' для выхода):")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if strings.ToLower(message) == "exit" {
			fmt.Println("Завершение соединения...")
			break
		}

		_, err = conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Ошибка при отправке:", err)
			break
		}

		reply := make([]byte, 1024)
		n, err := conn.Read(reply)
		if err != nil {
			fmt.Println("Ошибка при получении ответа:", err)
			break
		}

		fmt.Printf("Ответ от сервера: %s\n", string(reply[:n]))
	}
}
