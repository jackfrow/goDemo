package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		receivedData := string(buffer[:n])
		fmt.Println("Received data:", receivedData)

		// Echo back the received data
		_, err = conn.Write([]byte("Server echoed: " + receivedData))
		if err != nil {
			fmt.Println("Error writing:", err)
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is running. Press Ctrl+C to exit.")

	// Handle Ctrl+C to gracefully shut down the server
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)
	go func() {
		<-sigCh
		fmt.Println("\nShutting down the server...")
		listener.Close()
		os.Exit(0)
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			break
		}
		go handleConnection(conn)
	}
}
