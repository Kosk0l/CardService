package storage
// Storage

import (
	"CardService/proto/grpcProto"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

//===================================================================================================================//

// TODO: СДЕЛАТЬ В CLEAN ARHITECTURE

func (p *Postgres) CreateCardPG(ctx context.Context, req *grpcProto.CreateCardRequest) (*grpcProto.CardResponse, error) {
	query := `
		INSERT INTO cards (user_id, deck_id, text1, text2) VALUES
		($1, $2, $3, $4)
		RETURNING card_id;
	`
	var cardid int64
	err := p.pool.QueryRow(ctx, query, req.Userid, req.Deckid, req.Text1, req.Text2).Scan(&cardid)
	if err != nil {
		return nil, fmt.Errorf("failed to сreate new card: %w", err)
	}
	
	return &grpcProto.CardResponse{
		Cardid: cardid,
		Deckid: req.Deckid,
		Userid: req.Userid,
		Text1: req.Text1,
		Text2: req.Text2,
	}, nil
}

func (p *Postgres) UpdadeCardPG(ctx context.Context, req *grpcProto.UpdateCardRequest) (*grpcProto.CardResponse, error) {
	query := `
		UPDATE cards
		SET text1 = $1, text2 = $2 
		WHERE card_id = $3
		RETURNING deck_id, user_id;
	`
	var (
		deckid int64
		userid int64
	)
	err := p.pool.QueryRow(ctx, query, req.Text1, req.Text2, req.Cardid).Scan(&deckid, &userid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("card not found with id %d", req.Cardid)
		}
		return nil, fmt.Errorf("failed to update the card: %w", err)
	}

	return &grpcProto.CardResponse{
		Cardid: req.Cardid,
		Userid: userid,
		Deckid: deckid,
		Text1: req.Text1,
		Text2: req.Text2,
	}, nil
}

func (p *Postgres) GetCardPG(ctx context.Context, req *grpcProto.GetCardRequest) (*grpcProto.CardResponse, error) {
	query := `
		SELECT deck_id, user_id, text1, text2 FROM cards WHERE card_id = $1;
	`

	var (
		deckid int64
		userid int64
		text1 string
		text2 string
	)
	err := p.pool.QueryRow(ctx, query, req.Cardid).Scan(&deckid, &userid, &text1, &text2)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("card not found with id %d", req.Cardid)
		}
		return nil, fmt.Errorf("failed to get the card: %w", err)
	}

	return &grpcProto.CardResponse{
		Cardid: req.Cardid,
		Deckid: deckid,
		Userid: userid,
		Text1: text1,
		Text2: text2,
	}, nil
}

func (p *Postgres) DeleteCardPG(ctx context.Context, req *grpcProto.DeleteCardRequest) (*grpcProto.DeleteCardResponse, error) {
	query := `
		DELETE FROM cards WHERE card_id = $1;
	`

	cmd, err := p.pool.Exec(ctx, query, req.Cardid)
	if err != nil {
		return nil, fmt.Errorf("failed to get the card: %w", err)
	}
	// Проверка, что карточка действительно удалена
	if cmd.RowsAffected() == 0 {
		return nil, fmt.Errorf("card not found with id %d", req.Cardid)
	}

	return &grpcProto.DeleteCardResponse{
		Success: "success",
	}, nil
}

//===================================================================================================================//

func (p *Postgres) ListCardsPG() () {

}