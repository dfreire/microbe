package auth

type store interface {
	hasUserWithEmail(email string) (bool, error)
}
