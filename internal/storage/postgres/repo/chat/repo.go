package chat

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/postgres"
	"github.com/jmoiron/sqlx"
)

type Chat struct {
	ChatID    int64     `db:"chat_id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

type ChatDTO struct {
	ChatID     int64     `db:"chat_id"`
	Name       string    `db:"name"`
	CreatedAt  time.Time `db:"created_at"`
	Registered int64     `db:"registered"`
}

func (c *Chat) Add(ctx context.Context, tx *sqlx.Tx) error {
	query := `
		INSERT INTO "chat" (
			"chat_id",
			"name",
			"created_at"
		) VALUES (:chat_id, :name, :created_at)
		ON CONFLICT  ("chat_id")
		DO UPDATE SET "name" = :name;`

	_, err := tx.NamedExecContext(ctx, query, c)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func GetAll(ctx context.Context) ([]*ChatDTO, error) {
	var data []*ChatDTO

	query := `
		SELECT c.*, COUNT(uc.chat_id) AS registered 
		FROM "chat" AS c 
		INNER JOIN "user_chat" AS uc 
		    ON c.chat_id = uc.chat_id 
		GROUP BY uc.chat_id, c.chat_id, c.name, c.created_at;`

	err := postgres.GetDB().SelectContext(ctx, &data, query)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetByUser(ctx context.Context, value string) ([]*ChatDTO, error) {
	var data []*ChatDTO

	query := `
		SELECT c.*, COUNT(uc.chat_id) AS registered 
		FROM "chat" AS c 
		INNER JOIN "user_chat" AS uc 
		    ON c.chat_id = uc.chat_id 
		INNER JOIN "user" AS u 
			ON u.user_id = uc.user_id
		WHERE u.user_id = $1 
		   OR LOWER(u.username) LIKE $2 
		   OR LOWER(u.first_name) LIKE $2 
		   OR LOWER(u.last_name) LIKE $2
		GROUP BY uc.chat_id, c.chat_id, c.name, c.created_at
		ORDER BY registered DESC;`

	id, _ := strconv.ParseInt(value, 10, 64)
	v := strings.ToLower(value)

	err := postgres.GetDB().SelectContext(ctx, &data, query, id, v)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetByPeriod(ctx context.Context, from time.Time, to time.Time) ([]*Chat, error) {
	var data []*Chat

	query := `
		SELECT *
		FROM "chat"
		WHERE "created_at" >= $1 AND "created_at" < $2`

	err := postgres.GetDB().SelectContext(ctx, &data, query, from, to)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func HasByID(ctx context.Context, id int64) (bool, error) {
	var has bool

	query := `
		SELECT EXISTS (
		    SELECT 1 
		    FROM "chat" 
		    WHERE "chat_id" = $1
		)`

	err := postgres.GetDB().GetContext(ctx, &has, query, id)
	if err != nil {
		return false, err
	}

	return has, nil
}
