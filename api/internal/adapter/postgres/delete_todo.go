package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/domain"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/transaction"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (p *Postgres) DeleteTodo(ctx context.Context, id uuid.UUID) error{
	const sql = `UPDATE todo SET deleted_at = NOW() WHERE id = $1`

	txOrPool := transaction.TryExtractTX(ctx)

	_, err:= txOrPool.Exec(ctx, sql, id)
	if err != nil{
		if errors.Is(err, pgx.ErrNoRows){
			return domain.ErrNotFound
		}

		return fmt.Errorf("txOrPool.Exec: %w", err)
	}
	
	return nil
}