package api

import (
	"encoding/json"
	"io"
	"net/http"

	users "github.com/kyomel/go-scalable-educative/user"
)

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var request UpdateUserRequest
	err = json.Unmarshal(data, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := users.FindUserByName(request.CurrentName)
	if user != nil {
		err := user.UpdateUser(request.NewName, request.Age, request.CurrentVersion)
		if err != nil {
			resp, _ := json.Marshal(ErrorResponse{
				Error: err.Error(),
			})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(resp)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	resp, _ := json.Marshal(ErrorResponse{
		Error: "Unable to find user to update",
	})
	w.Write(resp)
}
