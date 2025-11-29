package services

// services

import (
	"CardService/internal/models"
	"context"
)

// Паттер Инверсии зависимостей
// Абстракция, принадлежащая высокому уровню 
type CardRepository interface {
	CreateCard(ctx context.Context, card models.Card) (int64, error)
    GetCard(ctx context.Context, cardid int64) (models.Card, error)
    UpdateCard(ctx context.Context, card models.Card) error
    DeleteCard(ctx context.Context, cardid int64) error
}

// Бизнес-Логика хендлеров
type CardService struct {
	repo CardRepository // Логика зависит только от интерфейса
}

//===================================================================================================================//

// Создание объекта
func NewService(repo CardRepository) *CardService {
	return &CardService{
		repo: repo,
	}
}

func (c *CardService) GetCardRepo(ctx context.Context, req int64) ( error) {

	return nil
}

func (c *CardService) CreateCardRepo() () {

}

func (c *CardService) UpdadeCardRepo() () {

}

func (c *CardService) DeleteCard() () {

}