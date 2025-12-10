package server

import (
	"CardService/proto/grpcProto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	grpc *grpc.Server
	port string
}

func NewGrpcServer(port string, handler grpcProto.CardServiceServer) (*GRPCServer) {
	s := grpc.NewServer()
	grpcProto.RegisterCardServiceServer(s, handler)

	return &GRPCServer{
		grpc: s,
		port: port,
	}
}

func (s *GRPCServer) Start() (error) {
	lis, err := net.Listen("tcp",s.port)
	if err != nil {
		return err
	}

	log.Println("Server started on port: ", s.port)
	return s.grpc.Serve(lis)
}

func (s *GRPCServer) Stop() () {
	log.Println("stopping gRPC server")
	s.grpc.GracefulStop()
}