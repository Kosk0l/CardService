package services

import "context"

type CardsService struct {
	getCard GetCard
	createCard CreateCard
	updateCard UpdateCard
	deleteCard DeleteCard
}

type GetCard interface {
	GetCard(ctx context.Context, ) ( , error)
}

type CreateCard interface {
	CreateCard(ctx context.Context, ) ( , error)
}

type UpdateCard interface {
	UpdateCard(ctx context.Context, ) ( , error)
}

type DeleteCard interface {
	DeleteCard(ctx context.Context, ) ( , error)
}