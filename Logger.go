package main_logger

import (
	"fmt"
	"net"
	"os"
)

func SendUDPLogs() {
	// Define the receiver's address
	receiverAddr := "127.0.0.1:8080" // Change as needed
	conn, err := net.Dial("udp", receiverAddr)
	if err != nil {
		fmt.Println("Error connecting to UDP server:", err)
		return
	}
	defer conn.Close()

	// Logs to send
	logs := []string{
		"Log 1: Application started",
		"Log 2: User logged in",
		"Log 3: Error occurred",
	}

	for _, log := range logs {
		// Send each log
		_, err := conn.Write([]byte(log))
		if err != nil {
			fmt.Println("Error sending log:", err)
			os.Exit(1)
		}
		fmt.Println("Sent:", log)
	}
}
