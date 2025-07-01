package api

import (
	"encoding/json"
	"io"
	"net/http"

	users "github.com/kyomel/go-scalable-educative/user"
)

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var request AddUserRequest
	err = json.Unmarshal(data, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users.AddUser(request.Name, request.Age)
	w.WriteHeader(http.StatusOK)
}
