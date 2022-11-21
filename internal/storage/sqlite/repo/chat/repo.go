package chat

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Chat struct {
	ChatID int64  `db:"chat_id"`
	Name   string `db:"name"`
}

func (c *Chat) Add(ctx context.Context, tx *sqlx.Tx) error {
	query := `
		INSERT INTO "chat" (
			"chat_id",
			"name"
		) VALUES (:chat_id, :name)
		ON CONFLICT  ("chat_id")
		DO UPDATE SET "name" = :name;`

	_, err := tx.NamedExecContext(ctx, query, c)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}
