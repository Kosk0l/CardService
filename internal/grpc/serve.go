package grpc

import (
	"CardService/proto/grpcProto"
	"context"
)

type Server struct {
	grpcProto.UnimplementedCardServiceServer // наследует генерированные данные
}

//===================================================================================================================//

func (s *Server) GetCard(ctx context.Context, req *grpcProto.GetCardRequest) (*grpcProto.CardResponse, error) {

	return nil ,nil
}

func (s *Server) CreateCard(ctx context.Context, req *grpcProto.CreateCardRequest) (*grpcProto.CardResponse, error) {

	return nil, nil	
}

func (s *Server) UpdateCard(ctx context.Context, req *grpcProto.UpdateCardRequest) (*grpcProto.CardResponse, error) {

	return nil, nil
}

func (s *Server) DeleteCard(ctx context.Context, req *grpcProto.DeleteCardRequest) (*grpcProto.DeleteCardResponse, error) {

	return nil, nil
}

func (s *Server) ListCard(req *grpcProto.ListCardRequest, stream grpcProto.CardService_ListCardServer) error {

	return nil
}
