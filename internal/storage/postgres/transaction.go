package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const failSerializeTransaction string = "SerializeTransaction() failed: %v"

func SerializeTransaction(ctx context.Context) (*sqlx.Tx, error) {
	tx, err := GetDB().BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return nil, fmt.Errorf(failSerializeTransaction, err)
	}

	return tx, nil
}
