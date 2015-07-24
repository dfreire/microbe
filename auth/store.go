package auth

import (
	"errors"

	"github.com/satori/go.uuid"
)

type store interface {
	getUser(domain, email string) (User, error)

	createUser(user User) error
	updateUser(user User) error
}

type implStoreMem struct {
	usersById map[interface{}]User
}

func NewStoreMem() store {
	return &implStoreMem{}
}

func (self implStoreMem) getUser(domain, email string) (User, error) {
	var user User
	for _, existentUser := range self.usersById {
		if existentUser.Domain == domain && existentUser.Email == email {
			user = existentUser
			break
		}
	}
	if user.Id == "" {
		return user, errors.New(ErrEntityNotFound)
	}
	return user, nil
}

func (self implStoreMem) createUser(user User) error {
	user.Id = uuid.NewV1()
	self.usersById[user.Id] = user
	return nil
}

func (self implStoreMem) updateUser(user User) error {
	self.usersById[user.Id] = user
	return nil
}
