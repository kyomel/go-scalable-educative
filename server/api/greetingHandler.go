package api

import (
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var welcomeMsg = "Welcome to the graceful server! 💃🏼\n"

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	greeting := welcomeMsg
	if name != "" {
		greeting = "Hello " + name + "!\n" + welcomeMsg
	} else {
		name = r.URL.Query().Get("name")
		if name != "" {
			greeting = "Hello " + name + "!\n" + welcomeMsg
		}
	}

	io.WriteString(w, greeting)
}
