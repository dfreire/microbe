package auth

type store interface {
	createUser(user User) error
	hasUserWithEmail(email string) (bool, error)
	putPendingUser(email string, password string) error
	getPendingUser(email string) (PendingUser, error)
}
