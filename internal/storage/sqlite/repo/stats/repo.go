package stats

import (
	"context"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite"
	"github.com/jmoiron/sqlx"
)

type Stats struct {
	ID          string `db:"id"`
	UserID      int64  `db:"user_id"`
	ChatID      int64  `db:"chat_id"`
	DinahuCount int64  `db:"dinahu_count"`
}

type StatsDTO struct {
	ID          string `db:"id"`
	UserID      int64  `db:"user_id"`
	ChatID      int64  `db:"chat_id"`
	DinahuCount int64  `db:"dinahu_count"`
	Username    string `db:"username"`
}

func (s *Stats) Add(ctx context.Context, tx *sqlx.Tx) error {
	query := `
		INSERT INTO "stats" (
			"id",
			"user_id",
			"chat_id",
			"dinahu_count"
		) VALUES (:id, :user_id, :chat_id, :dinahu_count)
		ON CONFLICT ("user_id", "chat_id") 
		DO NOTHING;`

	_, err := tx.NamedExecContext(ctx, query, s)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func (s *Stats) Update(ctx context.Context, tx *sqlx.Tx) error {
	query := `
		UPDATE "stats"
		SET "dinahu_count" = "dinahu_count" + 1
		WHERE "user_id" = :user_id AND "chat_id" = :chat_id`

	_, err := tx.NamedExecContext(ctx, query, s)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func GetByChatID(ctx context.Context, chatID int64) ([]*StatsDTO, error) {
	var data []*StatsDTO

	query := `
		SELECT s.*, u.username
		FROM "stats" AS s
		INNER JOIN "user" AS u
			ON s.user_id = u.user_id
		INNER JOIN "user_chat" AS uc 
			ON s.user_id = uc.user_id
		WHERE s.chat_id = ?
		ORDER BY s.dinahu_count DESC`

	err := sqlite.GetDB().SelectContext(ctx, &data, query, chatID)
	if err != nil {
		return nil, err
	}

	return data, nil
}
