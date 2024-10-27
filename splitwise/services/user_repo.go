package services

import "errors"

var userMap = make(map[string]*User)

func add(user *User) (bool, error) {
	_, ok := userMap[user.GetUserID()]
	if ok {
		return false, errors.New("user already exists")
	}
	userMap[user.GetUserID()] = user
	return true, nil
}

func get(id string) *User {
	return userMap[id]
}

func remove(id string) {
	delete(userMap, id)
}
