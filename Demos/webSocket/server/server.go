package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	// 定时向客户端发送消息
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			err := conn.WriteMessage(websocket.TextMessage, []byte("Server: Ping"))
			if err != nil {
				fmt.Println("Error writing message:", err)
				break
			}
		}
	}()

	for {
		// 读取客户端发送的消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		fmt.Printf("Received message: %s\n", string(msg))
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	fmt.Println("WebSocket server running on ws://localhost:8080/ws")
	http.ListenAndServe(":8080", nil)
}
