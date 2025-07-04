package api

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kyomel/go-scalable-educative/network"
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

	router.Get("/delay", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		w.Write([]byte("Done!\n"))
	})

	router.Get("/timeout", func(w http.ResponseWriter, r *http.Request) {
		resp, err := network.NewClient().Get("http://localhost:8080/delay")
		if err != nil {
			fmt.Printf("error in API - %v\n", err)
			w.Write([]byte("error!\n"))
			w.WriteHeader(http.StatusGatewayTimeout)
			return
		}

		resVal, _ := io.ReadAll(resp.Body)
		w.Write(resVal)
	})

	return router
}
