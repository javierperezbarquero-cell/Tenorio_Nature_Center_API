package api

import (
	"context"
	"fmt"
	"rest/dto"
)

func (server *Server) ExecTx(ctx context.Context, fn func(*dto.Queries) error) error {
	tx, err := server.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := dto.New(tx)

	if err := fn(q); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("error en transacción: %v, error en rollback: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}