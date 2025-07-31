package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "SERVER: ", log.Ldate|log.Ltime|log.Lshortfile)

	if isPortInUse(":8080") {
		logger.Fatal("Port 8080 is already in use")
	}

	srv := server.New(logger)
	logger.Println("Starting server on :8080")
	if err := srv.Start(); err != nil {
		logger.Fatalf("Server failed: %v", err)
	}
}

func isPortInUse(addr string) bool {
	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
