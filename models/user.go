package models

import (
	"errors"
	"strings"
)

// User -
type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
	tokens   []string
}

// TODO: Refactor.
var tokens = make(map[string]*User)

// AddToken -
func AddToken(token string, user *User) {
	tokens[token] = user
}

// RemoveToken -
func RemoveToken(token string) {
	delete(tokens, token)
}

// FindByToken -
func FindByToken(token string) *User {
	user := tokens[token]
	return user
}

var users = make(map[string]bool)

// RegisterUser -
func RegisterUser(username, password string) (*User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty")
	} else if !isUsernameAvailable(username) {
		return nil, errors.New("The username isn't available")
	}

	users[username] = true
	user := User{
		Username: username,
		Password: password,
		tokens:   make([]string, 0),
	}
	return &user, nil
}

func isUsernameAvailable(username string) bool {
	return !users[username]
}
