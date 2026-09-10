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

func (r *TodoTepo) UpdateTodo(id uuid.UUID, todo *domain.Todo, ctx context.Context) (*domain.Todo, error) {
	const query = `
		UPDATE todos
		SET title = $1, description = $2, completed = $3
		WHERE id = $4
		RETURNING id, title, description, completed, created_at
	`

	var updatedTodo domain.Todo

	if err := r.pool.QueryRow(
		ctx, query,
		todo.Title, todo.Description, todo.Completed, id,
	).Scan(
		&updatedTodo.ID,
		&updatedTodo.Title,
		&updatedTodo.Description,
		&updatedTodo.Completed,
		&updatedTodo.CreatedAt,
	); err != nil {
		return nil, fmt.Errorf("update todo: %w", err)
	}

	return &updatedTodo, nil
}

func (r *TodoTepo) DeleteTodo(id uuid.UUID, ctx context.Context) error {
	const query = `
		DELETE FROM todos
		WHERE id = $1
	`

	result, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete todo: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("todo with id %s not found", id)
	}

	return nil
}
