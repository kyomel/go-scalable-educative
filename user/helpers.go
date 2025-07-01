package users

import (
	"fmt"
)

func GetUsers() []User {
	userData.RLock()
	defer userData.RUnlock()
	return userData.users
}

func GetVersion() int32 {
	return userData.version
}

func FindUserByName(name string) *User {
	userData.RLock()
	defer userData.RUnlock()
	for i, user := range userData.users {
		if user.Name == name {
			return &userData.users[i]
		}
	}
	return nil
}

func FindUserByAge(age int32) *User {
	userData.RLock()
	defer userData.RUnlock()
	for i, user := range userData.users {
		if user.Age == age {
			return &userData.users[i]
		}
	}
	return nil
}

func AddUser(name string, age int32) {
	newUser := User{
		Name: name,
		Age:  age,
	}
	userData.Lock()
	defer userData.Unlock()
	userData.users = append(userData.users, newUser)
	userData.version += 1
}

func (user *User) UpdateUser(name string, age int32, currentVersion int32) error {
	if currentVersion != userData.version {
		return fmt.Errorf("version mismatch, expected v%d but found v%d", userData.version, currentVersion)
	}

	if name == "" {
		name = user.Name
	}
	if age <= 0 {
		age = user.Age
	}

	if name == user.Name && age == user.Age {
		return nil
	}

	userData.Lock()
	defer userData.Unlock()
	user.Name = name
	user.Age = age
	userData.version += 1
	return nil
}
