package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/domain"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/transaction"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)


func (p *Postgres) GetTodo(ctx context.Context, id uuid.UUID) (domain.Todo, error) {
	const sql = `SELECT created_at, updated_at, deleted_at, title, description, completed FROM todo WHERE id = $1`

	todo:= domain.Todo{}

	dto := struct{
		CreatedAt pgtype.Timestamptz
		UpdatedAt pgtype.Timestamptz
		DeletedAt pgtype.Timestamptz
		Title pgtype.Text
		Description pgtype.Text
		Completed pgtype.Bool
	}{}

	dest := []any{
		&dto.CreatedAt,
		&dto.UpdatedAt,
		&dto.DeletedAt,
		&dto.Title,
		&dto.Description,
		&dto.Completed,
	}

	txOrPool := transaction.TryExtractTX(ctx)

	err:= txOrPool.QueryRow(ctx, sql, id).Scan(dest...)

	if err != nil{
		if errors.Is(err, pgx.ErrNoRows){
			return todo, domain.ErrNotFound
		}
		return todo, fmt.Errorf("txOrPool.QueryRow: %w", err)
	}

	todo.ID = id
	todo.CreatedAt = dto.CreatedAt.Time
	todo.UpdatedAt = dto.UpdatedAt.Time	
	todo.DeletedAt = dto.DeletedAt.Time
	todo.Title = dto.Title.String
	todo.Description = dto.Description.String
	todo.Completed = dto.Completed.Bool

	return todo, nil
}