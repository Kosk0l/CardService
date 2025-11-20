package storage
// Storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

//===================================================================================================================//

// Обратиться к БД для создания новой карточки
func (p *Postgres) CreateCardPG(ctx context.Context, userid int64, deckid int64, text1 string, text2 string) (int64, error) {
	query := `
		INSERT INTO cards (user_id, deck_id, text1, text2) VALUES
		($1, $2, $3, $4)
		RETURNING card_id;
	`
	var cardid int64
	err := p.pool.QueryRow(ctx, query, userid, deckid, text1, text2).Scan(&cardid)
	if err != nil {
		return cardid, fmt.Errorf("failed to сreate new card: %w", err)
	}
	
	return cardid, nil 
	// Вернем только id новой карточки, 
	// Остальную часть структуры ответа взять в слоях выше
}


// Обратиться к Бд собновлением карточки по cardid, возвращаем deckid, userid
func (p *Postgres) UpdadeCardPG(ctx context.Context, cardid int64, text1 string, text2 string) (int64, int64, error) {
	query := `
		UPDATE cards
		SET text1 = $1, text2 = $2 
		WHERE card_id = $3
		RETURNING user_id, deck_id;
	`
	var (
		userid int64
		deckid int64
	)
	err := p.pool.QueryRow(ctx, query, text1, text2, cardid).Scan(&userid, &deckid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, 0, fmt.Errorf("card not found with id %d", cardid)
		}
		return 0, 0, fmt.Errorf("failed to update the card: %w", err)
	}

	return userid, deckid, nil
	// Остальную часть структуры ответа взять в слоях выше
}


// Получить данные карточки по id
func (p *Postgres) GetCardPG(ctx context.Context, cardid int64) (int64, int64, string, string, error) {
	query := `
		SELECT deck_id, user_id, text1, text2 FROM cards WHERE card_id = $1;
	`

	var (
		deckid int64
		userid int64
		text1 string
		text2 string
	)
	err := p.pool.QueryRow(ctx, query, cardid).Scan(&deckid, &userid, &text1, &text2)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, 0, "", "", fmt.Errorf("card not found with id %d", cardid)
		}
		return 0,0, "", "", fmt.Errorf("failed to get the card: %w", err)
	}

	return userid, deckid, text1, text2, nil
	// Взять id в слоях выше
}

// Обратиться к Бд для удаления карточки по cardid, возвращаем string ответ
func (p *Postgres) DeleteCardPG(ctx context.Context, cardid int64) (string, error) {
	query := `
		DELETE FROM cards WHERE card_id = $1;
	`

	cmd, err := p.pool.Exec(ctx, query, cardid)
	if err != nil {
		return "", fmt.Errorf("failed to get the card: %w", err)
	}
	// Проверка, что карточка действительно удалена
	if cmd.RowsAffected() == 0 {
		return "", fmt.Errorf("card not found with id %d", cardid)
	}

	return "success", nil
}

//===================================================================================================================//

func (p *Postgres) ListCardsPG() () {

}