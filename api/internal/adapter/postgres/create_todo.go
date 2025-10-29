package postgres

import (
	"context"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/domain"
)

func (p *Postgres) CreateTodo(ctx context.Context, todo domain.Todo) error{
	const sql = `INSERT INTO profile (id, title, description, completed)
					VALUES ($1, $2, $3, $4)`

					return nil
}