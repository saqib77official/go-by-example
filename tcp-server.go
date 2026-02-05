package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== TCP Server ===")

	// Start TCP server
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer listener.Close()

	fmt.Println("TCP Server listening on :8081")
	fmt.Println("Use: telnet localhost 8081 or nc localhost 8081")
	fmt.Println("Press Ctrl+C to stop")

	// Accept connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v\n", err)
			continue
		}

		// Handle connection in goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	
	// Get client address
	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("New connection from %s\n", clientAddr)

	// Send welcome message
	welcome := "Welcome to TCP Server!\n"
	conn.Write([]byte(welcome))

	// Read from connection
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("Message from %s: %s\n", clientAddr, message)

		// Handle commands
		response := handleCommand(message)
		conn.Write([]byte(response + "\n"))

		// Exit on quit
		if strings.ToLower(message) == "quit" {
			break
		}
	}

	fmt.Printf("Connection closed from %s\n", clientAddr)
}

func handleCommand(message string) string {
	msg := strings.ToLower(strings.TrimSpace(message))
	
	switch msg {
	case "help":
		return "Available commands: help, time, date, echo <text>, status, quit"
	case "time":
		return fmt.Sprintf("Current time: %s", time.Now().Format("15:04:05"))
	case "date":
		return fmt.Sprintf("Current date: %s", time.Now().Format("2006-01-02"))
	case "status":
		return "Server is running and healthy"
	case "":
		return "Please enter a command (try 'help')"
	default:
		if strings.HasPrefix(msg, "echo ") {
			return strings.TrimPrefix(msg, "echo ")
		}
		return fmt.Sprintf("Unknown command: %s. Type 'help' for available commands.", message)
	}
}

// Alternative simple echo server
func startEchoServer() {
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer listener.Close()

	fmt.Println("Echo server on :8082")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go func(c net.Conn) {
			defer c.Close()
			scanner := bufio.NewScanner(c)
			for scanner.Scan() {
				text := scanner.Text()
				fmt.Printf("Echo: %s\n", text)
				c.Write([]byte(text + "\n"))
			}
		}(conn)
	}
}
