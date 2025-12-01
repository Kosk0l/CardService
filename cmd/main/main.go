package main

import (
	pb "CardService/proto/grpcProto"
	"log"
	"net"
	"google.golang.org/grpc"
	grpcServer "CardService/internal/grpchandler"
)

func main() {
	// Настройка слушателя
	lis, err := net.Listen("tcp", ":44045")
	if err != nil {
		log.Fatalf("Error listen server: %v", err)
	}

	// Иннициализация сервера
	s := grpc.NewServer()
	pb.RegisterCardServiceServer(s, &grpcServer.Server{})

	log.Println("Server working")
	// Настройка серве слушателя
	if err := s.Serve(lis); err != nil { 
		log.Fatalf("failed to serve: %v", err)
	}
}