package services

// services

import (
	"CardService/internal/models"
	"context"
)

// Паттер Инверсии зависимостей
// Абстракция, принадлежащая высокому уровню 
type CardRepository interface {
	CreateCard(ctx context.Context, card *models.Card) (int64, error)
    GetCard(ctx context.Context, cardid int64) (models.Card, error)
    UpdateCard(ctx context.Context, card *models.Card) error
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

// Получение карточки; Бизнес-Логика
func (c *CardService) GetCardService(ctx context.Context, req int64) (models.Card, error) {
	return c.repo.GetCard(ctx, req)
}

// Создание карточки; Бизнес-Логика
func (c *CardService) CreateCardService(ctx context.Context, card models.Card) (int64, error) {
	return c.repo.CreateCard(ctx, &card)
}

func (c *CardService) UpdateCardService(ctx context.Context, card models.Card) (error) {
	return c.repo.UpdateCard(ctx, &card)
}

func (c *CardService) DeleteCardService(ctx context.Context, req int64) (error) {
	return c.repo.DeleteCard(ctx, req)
}