package api

import (
	"io"
	"net/http"

	"github.com/kyomel/go-scalable-educative/services"
)

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	res, err := services.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resData, _ := io.ReadAll(res.Body)
	w.Write(resData)
}
