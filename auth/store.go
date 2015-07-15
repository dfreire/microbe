package auth

type store interface {
	existsUser(app, email string) (bool, error)
	createUser(user User) error

	findUnconfirmedUser(app, email string) (UnconfirmedUser, error)
	createUnconfirmedUser(unconfirmedUser UnconfirmedUser) error
	removeUnconfirmedUser(unconfirmedUser UnconfirmedUser) error
}
