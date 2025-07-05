package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kyomel/go-scalable-educative/middlewares"
)

func GetRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middlewares.WithRateLimit)
	router.Get("/", BaseHandler)
	router.Get("/greeting", GreetingHandler)
	router.Get("/greeting/{name}", GreetingHandler)
	router.Get("/users", FindUserHandler)
	router.Post("/users", AddUserHandler)
	router.Patch("/users", UpdateUserHandler)
	router.Get("/data", GetAllData)

	return router
}
