package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func GetRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", BaseHandler)
	router.Get("/greeting", GreetingHandler)
	router.Get("/greeting/{name}", GreetingHandler)
	router.Get("/users", FindUserHandler)
	router.Post("/users", AddUserHandler)
	router.Patch("/users", UpdateUserHandler)

	router.Get("/timeout", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		w.Write([]byte("Done!"))
	})

	return router
}
