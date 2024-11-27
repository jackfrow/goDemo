package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	go func() {
		for {
			select {
			case <-time.Tick(time.Second):
				err := c.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket Server!"))
				if err != nil {
					log.Println("write:", err)
					return
				}
			}
		}
	}()

	for {
		select {
		case <-interrupt:
			fmt.Println("Interrupt signal received, closing connection...")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-time.After(time.Second):
			}
			c.Close()
			return
		}
	}
}
