package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Ошибка апгрейда:", err)
		return
	}
	defer ws.Close()
	clients[ws] = true
	log.Println("Новый клиент подключён")
	for {
		var msg string
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Println("Ошибка чтения:", err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}
func handleMessages() {
	for {
		msg := <-broadcast
		log.Println("Отправка:", msg)
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println("Ошибка отправки:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	fmt.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
