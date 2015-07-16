package auth

type store interface {
	getUser(domain, email string) (User, error)

	createUser(user User) error
	updateUser(user User, fields []string) error
}
