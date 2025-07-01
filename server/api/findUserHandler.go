package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	users "github.com/kyomel/go-scalable-educative/user"
)

func FindUserHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	if name == "" && age == "" {
		ListUsersHandler(w, r)
		return
	}

	var user *users.User

	if name != "" {
		user = users.FindUserByName(name)
	} else if age != "" {
		age, err := strconv.Atoi(age)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid age number passed\n"))
			return
		}

		user = users.FindUserByAge(int32(age))
	}

	if user != nil {
		w.Write([]byte(fmt.Sprintf("Found user with name %s and age %d", user.Name, user.Age)))
		return
	}

	w.WriteHeader(http.StatusNotFound)
	resp, _ := json.Marshal(ErrorResponse{
		Error: "Unable to fun requested user",
	})
	w.Write(resp)
}
