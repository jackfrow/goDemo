package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func sendDataToServer(conn net.Conn) {
	rand.Seed(time.Now().UnixNano())

	for {
		select {
		case <-time.After(time.Second):
			randomData := rand.Intn(1000)
			fmt.Printf("Sending data to server: %d\n", randomData)

			_, err := conn.Write([]byte(fmt.Sprintf("%d", randomData)))
			if err != nil {
				fmt.Println("Error sending data:", err.Error())
				return
			}

			buffer := make([]byte, 1024)
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
				return
			}

			fmt.Println("Server response:", string(buffer[:n]))

		}
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	go sendDataToServer(conn) // 启动 Goroutine 发送数据

	// 等待 Ctrl+C 信号来结束程序
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)
	<-sigCh

	fmt.Println("Ctrl+C pressed. Exiting...")
}
