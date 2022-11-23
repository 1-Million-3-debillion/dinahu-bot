package info

import (
	"context"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite"
)

type Info struct {
	Chats int64 `db:"chats"`
	Users int64 `db:"users"`
}

func GetInfo(ctx context.Context, from int64, to int64) (*Info, error) {
	var model Info

	query := `
		SELECT t1.chats, t2.users FROM (
			SELECT COUNT(chat_id) AS chats 
			FROM chat AS c
			WHERE c.created_at >= ? AND c.created_at < ?
		) AS t1 
		INNER JOIN (
			SELECT COUNT(user_id) AS users 
			FROM user AS u
			WHERE u.created_at >= ? AND u.created_at < ?
		) AS t2`

	err := sqlite.GetDB().GetContext(ctx, &model, query, from, to, from, to)
	if err != nil {
		return nil, err
	}

	return &model, nil
}
