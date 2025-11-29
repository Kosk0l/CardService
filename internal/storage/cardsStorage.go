package storage

// Storage

import (
	"CardService/internal/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
)

//===================================================================================================================//

// Обратиться к БД для создания новой карточки
func (p *Postgres) CreateCard(ctx context.Context, card *models.Card) (int64, error) {
	query := `
		INSERT INTO cards (user_id, deck_id, text1, text2) VALUES
		($1, $2, $3, $4)
		RETURNING card_id;
	`
	var cardid int64
	err := p.pool.QueryRow(ctx, query, card.UserId, card.DeckId, card.Text1, card.Text2).Scan(&cardid)
	if err != nil {
		return cardid, fmt.Errorf("failed to сreate new card: %w", err)
	}
	
	return cardid, nil 
}

// Обратиться к Бд с обновлением карточки по cardid
func (p *Postgres) UpdateCard(ctx context.Context, card *models.Card) error {
	query := `
		UPDATE cards
		SET text1 = $1, text2 = $2
		WHERE card_id = $3;
	`

	cmd, err := p.pool.Exec(ctx, query, card.Text1, card.Text2, card.CardId)
	if err != nil {
		return fmt.Errorf("failed to update card: %w", err)
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("card not found with id %d", card.CardId)
	}

	return nil
}

// Получить данные карточки по id
func (p *Postgres) GetCard(ctx context.Context, cardid int64) (models.Card, error) {
	query := `
		SELECT card_id, user_id, deck_id, text1, text2
		FROM cards WHERE card_id = $1;
	`

	var card models.Card
	err := p.pool.QueryRow(ctx, query, cardid).Scan(&card.CardId, &card.UserId, &card.DeckId, &card.Text1, &card.Text2)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Card{}, fmt.Errorf("card not found with id %d", cardid)
		}
		return models.Card{}, fmt.Errorf("failed to get card: %w", err)
	}

	return card, nil
}

// Обратиться к Бд для удаления карточки по cardid, возвращаем string ответ
func (p *Postgres) DeleteCard(ctx context.Context, cardid int64) (error) {
	query := `
		DELETE FROM cards WHERE card_id = $1;
	`

	cmd, err := p.pool.Exec(ctx, query, cardid)
	if err != nil {
		return  fmt.Errorf("failed to get the card: %w", err)
	}
	// Проверка, что карточка действительно удалена
	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("card not found with id %d", cardid)
	}

	return nil
}

//===================================================================================================================//

func (p *Postgres) ListCardsPG() () {

}