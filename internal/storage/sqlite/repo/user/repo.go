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
}

func (u *User) Add(ctx context.Context, tx *sqlx.Tx) error {
	query := `
		INSERT INTO "user" (
			"user_id",
			"first_name",
			"last_name",
		    "username"
		) VALUES (:user_id, :first_name, :last_name, :username)
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

func GetByID(ctx context.Context, id int64) (*User, error) {
	var model User

	query := `
		SELECT *
		FROM "user"
		WHERE "user_id" = $1`

	err := sqlite.GetDB().GetContext(ctx, &model, query, id)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func GetByChatID(ctx context.Context, chatID int64) ([]*User, error) {
	var data []*User

	query := `
		SELECT u.*
		FROM "user" AS u
		INNER JOIN "user_chat" AS uc
			ON u.user_id = uc.user_id
		WHERE uc.chat_id = $1`

	err := sqlite.GetDB().SelectContext(ctx, &data, query, chatID)
	if err != nil {
		return nil, err
	}

	return data, nil
}
