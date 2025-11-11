package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Настройка слушателя
	lis, err := net.Listen("tcp", "44046")
	if err != nil {
		log.Fatalf("Error listen server: %v", err)
	}

	// Иннициализация сервера
	s := grpc.NewServer()

	// Настройка серве слушателя
	if err := s.Serve(lis); err != nil { 
		log.Fatalf("failed to serve: %v", err)
	}
}