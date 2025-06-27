package server

type User struct {
	name string
	age  int
}

var users = []User{
	{name: "John", age: 30},
	{name: "Jane", age: 25},
	{name: "Bob", age: 35},
}

func getUsers() []User {
	return users
}
