package repository

import (
	"context"
	"todo-list/internal/domain"

	"github.com/google/uuid"
)

type TodoRepo interface {
	AddTodo(todo *domain.Todo, ctx context.Context) (*domain.Todo, error)
	GetTodo(id uuid.UUID, ctx context.Context) (*domain.Todo, error)
	GetAllTodos(ctx context.Context) ([]domain.Todo, error)
	UpdateTodo(id uuid.UUID, todo *domain.Todo, ctx context.Context) (*domain.Todo, error)
	DeleteTodo(id uuid.UUID, ctx context.Context) error
}
