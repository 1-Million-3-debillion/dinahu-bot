package user

import (
	"context"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite"
	"github.com/jmoiron/sqlx"
)

type User struct {
	UserID    int64  `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Username  string `db:"username"`
	CreatedAt int64  `db:"created_at"`
}

type UserDTO struct {
	UserID    int64  `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Username  string `db:"username"`
	CreatedAt int64  `db:"created_at"`
	ChatName  string `db:"chat_name"`
}

func (u *User) Add(ctx context.Context, tx *sqlx.Tx) error {
	query := `
		INSERT INTO "user" (
			"user_id",
			"first_name",
			"last_name",
		    "username",
			"created_at"
		) VALUES (:user_id, :first_name, :last_name, :username, :created_at)
		ON CONFLICT ("user_id")
		DO UPDATE SET "first_name" = :first_name,
		              "last_name" = :last_name,
					  "username" = :username;`

	_, err := tx.NamedExecContext(ctx, query, u)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func (u *User) Delete(ctx context.Context, tx *sqlx.Tx) error {
	query := `
		DELETE FROM "user"
		WHERE "user_id" = :user_id`

	_, err := tx.NamedExecContext(ctx, query, u)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func GetByChatID(ctx context.Context, chatID int64) ([]*UserDTO, error) {
	var data []*UserDTO

	query := `
		SELECT u.*, c.name AS chat_name
		FROM "user" AS u
		INNER JOIN "user_chat" AS uc
			ON u.user_id = uc.user_id
		INNER JOIN "chat" AS c 
			ON uc.chat_id = c.chat_id
		WHERE uc.chat_id = $1`

	err := sqlite.GetDB().SelectContext(ctx, &data, query, chatID)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetByPeriod(ctx context.Context, from int64, to int64) ([]*UserDTO, error) {
	var data []*UserDTO

	query := `
		SELECT u.*, c.name AS chat_name
		FROM "user" AS u
		INNER JOIN "user_chat" AS uc
			ON u.user_id = uc.user_id
		INNER JOIN "chat" AS c 
			ON uc.chat_id = c.chat_id
		WHERE u.created_at >= $1 AND u.created_at < $2`

	err := sqlite.GetDB().SelectContext(ctx, &data, query, from, to)
	if err != nil {
		return nil, err
	}

	return data, nil
}
