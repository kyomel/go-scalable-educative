package api

import users "github.com/kyomel/go-scalable-educative/user"

type AddUserRequest struct {
	Name           string `json:"name"`
	Age            int32  `json:"age"`
	CurrentVersion int32  `json:"current_version"`
}

type UpdateUserRequest struct {
	CurrentName    string `json:"current_name"`
	NewName        string `json:"new_name"`
	Age            int32  `json:"age"`
	CurrentVersion int32  `json:"current_version"`
}

type ListUsersResponse struct {
	Users   []users.User `json:"users"`
	Version int32        `json:"current_version"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
