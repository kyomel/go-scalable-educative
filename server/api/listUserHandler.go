package api

import (
	"io"
	"net/http"

	"github.com/kyomel/go-scalable-educative/network"
)

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	res, err := network.NewClient().
		Name("fakerapi").
		Timeout(10).
		Get("https://fakerapi.it/api/v1/persons")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resData, _ := io.ReadAll(res.Body)
	w.Write(resData)
}
