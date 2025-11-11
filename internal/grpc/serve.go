package grpc

import (
	"CardService/proto/grpcProto"
	"context"
)

type Server struct {
	grpcProto.UnimplementedCardServiceServer // наследует генерированные данные
}

//===================================================================================================================//

func (s *Server) GetCardRequest(ctx context.Context, req *grpcProto.GetCardRequest) (*grpcProto.CardResponse, error) {

	return nil ,nil
}

func (s *Server) CreateCardRequest(ctx context.Context, req *grpcProto.CreateCardRequest) (*grpcProto.CardResponse, error) {

	return nil, nil	
}

func (s *Server) UpdateCardRequest(ctx context.Context, req *grpcProto.UpdateCardRequest) (*grpcProto.CardResponse, error) {

	return nil, nil
}

func (s *Server) DeleteCardRequest(ctx context.Context, req *grpcProto.CreateCardRequest) (*grpcProto.DeleteCardResponse, error) {

	return nil, nil
}

func (s *Server) ListCardRequest(ctx context.Context, req *grpcProto.ListCardRequest) (*grpcProto.CardResponse, error) {

	return nil, nil
}
