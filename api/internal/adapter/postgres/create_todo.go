package postgres

import (
	"context"
	"fmt"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/domain"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/transaction"
)

func (p *Postgres) CreateTodo(ctx context.Context, todo domain.Todo) error{
	const sql = `INSERT INTO todo (id, title, description, completed) VALUES ($1, $2, $3, $4)`

	args :=[]any{
		todo.ID,
		todo.Title,
		todo.Description,
		todo.Completed,
	}

	txOrPool := transaction.TryExtractTX(ctx)

	_, err := txOrPool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("txOrPool.Exec: %w", err)
	}

	return nil
}