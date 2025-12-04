package grpchandler

// Handlers - grpc

import (
	"CardService/internal/models"
	"CardService/internal/services"
	"CardService/proto/grpcProto"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Структура Handlers
type Server struct {
	grpcProto.UnimplementedCardServiceServer // наследует генерированные данные PROTOC
	service *services.CardService // Объект структуры Бизнес-логики (service)
}

// Конструктор сервера
func NewServer(service *services.CardService) *Server {
	return &Server{
		service: service,
	}
}

// Для преобразования модели в proto
func cardToProto(card models.Card) *grpcProto.CardResponse {
	return &grpcProto.CardResponse{
		Cardid: card.CardId,
		Userid: card.UserId,
		Deckid: card.DeckId,
		Text1:  card.Text1,
		Text2:  card.Text2,
	}
}

//===================================================================================================================//

// Получить карточку
func (s *Server) GetCard(ctx context.Context, req *grpcProto.GetCardRequest) (*grpcProto.CardResponse, error) {
	if req.Cardid == 0 {
		return nil, status.Error(codes.InvalidArgument, "error cardid")
	}
	
	// Запуск бизнес-логики
	card, err := s.service.GetCardService(ctx, req.Cardid)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "error in service")
	}

	return cardToProto(card), nil
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

	card := models.Card{
		UserId: req.Userid,
		DeckId: req.Deckid,
		Text1: req.Text1,
		Text2: req.Text2,
	}

	//запуск бизнес-логики
	id, err := s.service.CreateCardService(ctx, card)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "error in service")
	} 
	card.CardId = id

	return cardToProto(card), nil	
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

	card := models.Card{
		CardId: req.Cardid,
		Text1: req.Text1,
		Text2: req.Text2,
	}

	//запуск бизнес-логики
	err := s.service.UpdateCardService(ctx, card)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "error in service")
	}

	return cardToProto(card), nil
}

// Удалить карточку
func (s *Server) DeleteCard(ctx context.Context, req *grpcProto.DeleteCardRequest) (*grpcProto.DeleteCardResponse, error) {

	if req.Cardid == 0 {
		return nil, status.Error(codes.InvalidArgument, "error cardid")
	}

	// запуск бизнес-логики
	err := s.service.DeleteCardService(ctx, req.Cardid)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "error in service")
	}

	return &grpcProto.DeleteCardResponse{
		Success: "Success",
	}, nil
}

//===================================================================================================================//

// Стрим Карточек
func (s *Server) ListCard(req *grpcProto.ListCardRequest, stream grpcProto.CardService_ListCardServer) error {

	//TODO: запуск бизнес-логики

	return nil
}
