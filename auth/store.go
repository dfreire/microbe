package auth

type store interface {
	getUser(domain, email string) (User, error)

	createUser(user User) error
	updateUser(user User) error
}

type implStoreMem struct {
	usersById map[string]User
}

func NewStoreMem() store {
	return &implStoreMem{}
}

func (self implStoreMem) getUser(domain, email string) (User, error) {
	var user User
	return user, nil
}

func (self implStoreMem) createUser(user User) error {
	return nil
}

func (self implStoreMem) updateUser(user User) error {
	return nil
}
