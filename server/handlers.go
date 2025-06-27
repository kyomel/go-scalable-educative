package server

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var helloMsg = "Hello student!\n"
var welcomeMsg = "Welcome to the graceful server! 💃🏼\n"

func baseHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, helloMsg)
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
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

func findUser(w http.ResponseWriter, r *http.Request) {
	users := getUsers()
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	if name == "" && age == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No request params passed\n"))
		return
	} else if name != "" {
		for _, user := range users {
			if user.name == name {
				msg := fmt.Sprintf("Found user with name %s and age %d!\n", name, user.age)
				io.WriteString(w, msg)
				w.WriteHeader(http.StatusOK)
				return
			}
		}
	} else if age != "" {
		age, err := strconv.Atoi(age)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid age number passed\n"))
			return
		}
		for _, user := range users {
			if user.age == age {
				msg := fmt.Sprintf("Found user with name %s and age %d!\n", user.name, age)
				io.WriteString(w, msg)
				w.WriteHeader(http.StatusOK)
				return
			}
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Unable to find requested user\n"))
}
