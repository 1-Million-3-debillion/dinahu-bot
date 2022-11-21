package userChat

import (
	"context"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite"
	"github.com/jmoiron/sqlx"
)

type UserChat struct {
	ID          string `db:"id"`
	UserID      int64  `db:"user_id"`
	ChatID      int64  `db:"chat_id"`
	DinahuCount int64  `db:"dinahu_count"`
}

type UserChatDTO struct {
	ID          string `db:"id"`
	UserID      int64  `db:"user_id"`
	ChatID      int64  `db:"chat_id"`
	DinahuCount int64  `db:"dinahu_count"`
	Username    string `db:"username"`
}

func (uc *UserChat) Add(ctx context.Context, tx *sqlx.Tx) error {
	query := `
		INSERT INTO "user_chat" (
			"id",
			"user_id",
			"chat_id",
		    "dinahu_count"
		) VALUES (:id, :user_id, :chat_id, 0)
		ON CONFLICT ("user_id", "chat_id") 
		DO NOTHING;`

	_, err := tx.NamedExecContext(ctx, query, uc)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func (uc *UserChat) AddDinahu(ctx context.Context, tx *sqlx.Tx) error {
	query := `
		UPDATE "user_chat"
		SET "dinahu_count" = "dinahu_count" + 1
		WHERE "user_id" = :user_id AND "chat_id" = :chat_id;`

	_, err := tx.NamedExecContext(ctx, query, uc)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func GetByChatID(ctx context.Context, chatID int64) ([]*UserChatDTO, error) {
	var data []*UserChatDTO

	query := `
		SELECT uc.*, u.username
		FROM "user_chat" AS uc
		INNER JOIN "user" AS u 
			ON uc.user_id = u.user_id
		WHERE uc.chat_id = ?
		ORDER BY uc.dinahu_count DESC`

	err := sqlite.GetDB().SelectContext(ctx, &data, query, chatID)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func HasUserInChat(ctx context.Context, userID int64, chatID int64) (bool, error) {
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
