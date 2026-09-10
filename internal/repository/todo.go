package repository

import (
	"todo-list/internal/domain"

	"github.com/google/uuid"
)

type TodoRepo interface {
	AddTodo(todo *domain.Todo) error
	GetTodo(id uuid.UUID) (*domain.Todo, error)
	GetAllTodos() ([]domain.Todo, error)
	UpdateTodo(todo *domain.Todo) error
	DeleteTodo(id uuid.UUID) error
}
