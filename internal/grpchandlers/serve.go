package grpchandlers

// Handlers - grpc

import (
	"CardService/proto/grpcProto"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	grpcProto.UnimplementedCardServiceServer // наследует генерированные данные
}

//===================================================================================================================//

// Получить карточку
func (s *Server) GetCard(ctx context.Context, req *grpcProto.GetCardRequest) (*grpcProto.CardResponse, error) {
	if req.Cardid == 0 {
		return nil, status.Error(codes.InvalidArgument, "error cardid")
	}

	// запуск бизнес-логики
	

	return nil ,nil
}

// Создать карточку
func (s *Server) CreateCard(ctx context.Context, req *grpcProto.CreateCardRequest) (*grpcProto.CardResponse, error) {
	if req.Userid == 0 {
		return nil, status.Error(codes.InvalidArgument, "error userid")
	}

	if req.Deckid == 0 {
		return nil, status.Error(codes.InvalidArgument, "error deckid")
	}

	if req.Text1 == "" {
		return nil, status.Error(codes.InvalidArgument, "error Text1")
	}

	if req.Text2 == "" {
		return nil, status.Error(codes.InvalidArgument, "error Text2")
	}

	//TODO: запуск бизнес-логики

	return nil, nil	
}

// Обновить карточку
func (s *Server) UpdateCard(ctx context.Context, req *grpcProto.UpdateCardRequest) (*grpcProto.CardResponse, error) {
	if req.Cardid == 0 {
		return nil, status.Error(codes.InvalidArgument, "error cardid")
	}

	if req.Text1 == "" {
		return nil, status.Error(codes.InvalidArgument, "error text1")
	}

	if req.Text2 == "" {
		return nil, status.Error(codes.InvalidArgument, "error text2")
	}


	//TODO: запуск бизнес-логики

	return nil, nil
}

// Удалить карточку
func (s *Server) DeleteCard(ctx context.Context, req *grpcProto.DeleteCardRequest) (*grpcProto.DeleteCardResponse, error) {

	if req.Cardid == 0 {
		return nil, status.Error(codes.InvalidArgument, "error cardid")
	}

	//TODO: запуск бизнес-логики

	return nil, nil
}

//===================================================================================================================//

// Стрим Карточек
func (s *Server) ListCard(req *grpcProto.ListCardRequest, stream grpcProto.CardService_ListCardServer) error {

	//TODO: запуск бизнес-логики

	return nil
}
