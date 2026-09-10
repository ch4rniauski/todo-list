package service

import (
	"context"
	"todo-list/internal/domain"
	"todo-list/internal/repository"

	"github.com/google/uuid"
)

type TodoService struct {
	todoRepo repository.Todo
}

func NewTodoService(todoRepo repository.Todo) *TodoService {
	return &TodoService{todoRepo: todoRepo}
}

func (s *TodoService) CreateTodo(todo *domain.Todo, ctx context.Context) (*domain.Todo, error) {
	return s.todoRepo.AddTodo(todo, ctx)
}

func (s *TodoService) GetTodo(id uuid.UUID, ctx context.Context) (*domain.Todo, error) {
	return s.todoRepo.GetTodo(id, ctx)
}

func (s *TodoService) GetAllTodos(ctx context.Context) ([]domain.Todo, error) {
	return s.todoRepo.GetAllTodos(ctx)
}

func (s *TodoService) UpdateTodo(id uuid.UUID, todo *domain.Todo, ctx context.Context) (*domain.Todo, error) {
	return s.todoRepo.UpdateTodo(id, todo, ctx)
}

func (s *TodoService) DeleteTodo(id uuid.UUID, ctx context.Context) error {
	return s.todoRepo.DeleteTodo(id, ctx)
}
