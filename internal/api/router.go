package api

import (
	"todo-list/internal/api/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(todoHandler *handler.TodoHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/health", handler.Health)

	r.Group(func(r chi.Router) {
		r.Post("/todos", todoHandler.CreateTodo)
		r.Get("/todos/{id:uuid}", todoHandler.GetTodo)
		r.Get("/todos", todoHandler.GetAllTodos)
	})

	return r
}
