package postgres

import (
	"context"
	"fmt"
	"time"
	"todo-list/internal/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

	if err := r.pool.QueryRow(
		ctx, query,
		todo.ID, todo.Title, todo.Description, false, time.Now().UTC(),
	).Scan(
		&createdTodo.ID,
		&createdTodo.Title,
		&createdTodo.Description,
		&createdTodo.Completed,
		&createdTodo.CreatedAt,
	); err != nil {
		return nil, fmt.Errorf("add todo: %w", err)
	}

	return &createdTodo, nil
}

func (r *TodoTepo) GetTodo(id uuid.UUID, ctx context.Context) (*domain.Todo, error) {
	const query = `
		SELECT id, title, description, completed, created_at
		FROM todos t
		WHERE t.id = $1
	`

	var todo domain.Todo

	if err := r.pool.QueryRow(ctx, query, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Description,
		&todo.Completed,
		&todo.CreatedAt,
	); err != nil {
		return nil, fmt.Errorf("get todo by id: %w", err)
	}

	return &todo, nil
}

func (r *TodoTepo) GetAllTodos(ctx context.Context) ([]domain.Todo, error) {
	const query = `
		SELECT id, title, description, completed, created_at
		FROM todos
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("get all todos: %w", err)
	}
	defer rows.Close()

	todos, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Todo])
	if err != nil {
		return nil, fmt.Errorf("get all todos: %w", err)
	}

	return todos, nil
}
