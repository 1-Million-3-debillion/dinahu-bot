package info

import (
	"context"
	"time"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/postgres"
)

type Info struct {
	Chats int64 `db:"chats"`
	Users int64 `db:"users"`
}

func GetInfo(ctx context.Context, from time.Time, to time.Time) (*Info, error) {
	var model Info

	query := `
		SELECT t1.chats, t2.users FROM (
			SELECT COUNT("chat_id") AS chats 
			FROM "chat" AS c
			WHERE c.created_at >= $1 AND c.created_at < $2
		) AS t1 
		INNER JOIN (
			SELECT COUNT("user_id") AS users 
			FROM "user" AS u
			WHERE u.created_at >= $1 AND u.created_at < $2
		) AS t2
		ON t1.chats = t1.chats`

	err := postgres.GetDB().GetContext(ctx, &model, query, from, to)
	if err != nil {
		return nil, err
	}

	return &model, nil
}
