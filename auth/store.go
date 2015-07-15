package auth

type store interface {
	createUser(user User) error
	hasUserWithEmail(email string) (bool, error)

	getPendingUser(email string) (PendingUser, error)
	putPendingUser(email string, password string) error
	delPendingUser(email string) error
}
