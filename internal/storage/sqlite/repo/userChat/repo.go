package userChat

import (
	"context"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite"
	"github.com/jmoiron/sqlx"
)

type UserChat struct {
	ID     string `db:"id"`
	UserID int64  `db:"user_id"`
	ChatID int64  `db:"chat_id"`
}

func (uc *UserChat) Add(ctx context.Context, tx *sqlx.Tx) error {
	query := `
		INSERT INTO "user_chat" (
			"id",
			"user_id",
			"chat_id"
		) VALUES (:id, :user_id, :chat_id)
		ON CONFLICT ("user_id", "chat_id") 
		DO NOTHING;`

	_, err := tx.NamedExecContext(ctx, query, uc)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func (uc *UserChat) HasUserInChat(ctx context.Context, userID int64, chatID int64) (bool, error) {
	var has bool

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM "user_chat"
		    WHERE "user_id" = ? AND "chat_id" = ?
		);`

	err := sqlite.GetDB().GetContext(ctx, &has, query, userID, chatID)
	if err != nil {
		return false, err
	}

	return has, nil
}
