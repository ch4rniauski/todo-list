package postgres

import (
	"context"
	"fmt"
	"time"
	"todo-list/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoTepo struct {
	pool *pgxpool.Pool
}

func NewTodoRepo(pool *pgxpool.Pool) *TodoTepo {
	return &TodoTepo{pool: pool}
}

func (r *TodoTepo) AddTodo(todo *domain.Todo, ctx context.Context) (*domain.Todo, error) {
	const query = `
		INSERT INTO todos (id, title, description, completed, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, title, description, completed, created_at
	`

	var createdTodo domain.Todo

	err := r.pool.QueryRow(
		ctx, query,
		todo.ID, todo.Title, todo.Description, false, time.Now().UTC(),
	).Scan(
		&createdTodo.ID,
		&createdTodo.Title,
		&createdTodo.Description,
		&createdTodo.Completed,
		&createdTodo.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("add todo: %w", err)
	}

	return &createdTodo, nil
}
