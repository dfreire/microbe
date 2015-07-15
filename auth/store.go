package auth

type store interface {
	createUser(user User) error
	hasUserWithEmail(email string) (bool, error)

	getUnconfirmedUser(email string) (UnconfirmedUser, error)
	putUnconfirmedUser(email string, password string) error
	delUnconfirmedUser(email string) error
}
