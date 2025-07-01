package api

import (
	"encoding/json"
	"io"
	"net/http"

	users "github.com/kyomel/go-scalable-educative/user"
)

type fakeUsersAPIResponse struct {
	Data []users.FakeUser `json:"data"`
}

type apiUser struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://fakerapi.it/api/v1/persons")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if res.StatusCode != 200 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resData, _ := io.ReadAll(res.Body)
	var fakeUsers fakeUsersAPIResponse
	_ = json.Unmarshal(resData, &fakeUsers)

	var resp []apiUser
	for i := 0; i < len(fakeUsers.Data); i++ {
		resp = append(resp, apiUser{
			Name:   fakeUsers.Data[i].FirstName + " " + fakeUsers.Data[i].LastName,
			Gender: fakeUsers.Data[i].Gender,
		})
	}
	respBytes, _ := json.Marshal(resp)
	w.Write(respBytes)
}
