package users

import "sync"

type User struct {
	Name string `json:"name,omitempty"`
	Age  int32  `json:"age,omitempty"`
}

type UserData struct {
	sync.RWMutex
	users   []User
	version int32
}
