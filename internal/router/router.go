package router

import (
	"myTodo/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetUpRouter() http.Handler{
	r := chi.NewRouter()

	// r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)
	// r.Use(middleware.AllowContentType("application/json"))

	r.Route("/todos", func(r chi.Router){
		r.Get("/", handlers.GetTodos)
		r.Get("/{id}", handlers.GetTodo)
		r.Post("/", handlers.CreateTodo)
		r.Put("/{id}", handlers.UpdateTodo)
		r.Delete("/{id}", handlers.DeleteTodo)
	})

	return r
}
