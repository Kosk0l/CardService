package services

// services

import (
	"CardService/internal/models"
	"context"
)
type CardsService struct {
	getCard 	GetCard
	createCard 	CreateCard
	updateCard 	UpdateCard
	deleteCard 	DeleteCard
}

const deleteSuccess = "success" // Для Delete ручки


type GetCard interface {
	GetCard(ctx context.Context, req int64) (models.Card, error)
}

type CreateCard interface {
	CreateCard(ctx context.Context,  ) (models.Card, error)
}

type UpdateCard interface {
	UpdateCard(ctx context.Context, ) (models.Card,  error)
}

type DeleteCard interface {
	DeleteCard(ctx context.Context, req int64) (string, error)
}

//===================================================================================================================//

func (c *CardsService) GetCard(ctx context.Context, req int64) (models.Card, error) {

	return models.Card{
		
	}, nil
}

func (c *CardsService) CreateCard() () {

}

func (c *CardsService) UpdadeCard() () {

}

func (c *CardsService) DeleteCard() () {

}