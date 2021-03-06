package auth

import (
	"errors"
	"time"
)

type Auth interface {
	SignUp(domain, email string, password []byte) (signUpToken string, err error)
	ConfirmSignUp(signUpToken string) error

	// RemoveUnconfirmedUsers(elapsedTime time.Duration) error

	RequestResetPasswordToken(domain, email string) (resetPasswordToken string, err error)
	ResetPassword(resetPasswordToken string, newPassword []byte) error

	SignIn(domain, email string, password []byte) (sessionToken string, err error)
	ChangeEmail(sessionToken string, password []byte, newEmail string) error
	ChangePassword(sessionToken string, oldPassword []byte, newPassword []byte) error

	//GetAllUsers(adminToken string) ([]User, error)
	//GetUser(domain, email string) (User, error)
	//CreateUsers(adminToken, users []User) error
	//UpdateUsers(adminToken, users []User) error
	//RemoveUsers(adminToken, users []User) error
}

type implAuth struct {
	store      store
	adminToken string
}

func New(store store, adminToken string) Auth {
	return &implAuth{
		store:      store,
		adminToken: adminToken,
	}
}

func (self implAuth) SignUp(domain, email string, password []byte) (signUpToken string, err error) {
	user, err := self.store.getUser(domain, email)
	if err != nil {
		return "", err
	}

	if user.Id == nil {
		// new user
		user = User{
			CreatedAt:   time.Now(),
			Domain:      domain,
			Email:       email,
			password:    password,
			isConfirmed: false,
		}
		if err = self.store.createUser(user); err != nil {
			return "", err
		}

	} else {
		// user exists
		if user.Domain == domain && user.Email == email && !user.isConfirmed {
			// an unconfirmed user is trying to sign up again
			// do nothing, so a valid signUpToken will be returned
		} else {
			return "", errors.New(ErrEntityAlreadyExists)
		}
	}

	return createSignUpToken(domain, email)
}

func (self implAuth) ConfirmSignUp(signUpToken string) error {
	domain, email, err := parseSignUpToken(signUpToken)
	if err != nil {
		return err
	}

	user, err := self.store.getUser(domain, email)
	if err != nil {
		return err
	}

	user.isConfirmed = true
	if err = self.store.updateUser(user); err != nil {
		return err
	}

	return nil
}

func (self implAuth) RequestResetPasswordToken(domain, email string) (resetPasswordToken string, err error) {
	return "", nil
}

func (self implAuth) ResetPassword(resetPasswordToken string, newPassword []byte) error {
	return nil
}

func (self implAuth) SignIn(domain, email string, password []byte) (sessionToken string, err error) {
	return "", nil
}

func (self implAuth) ChangeEmail(sessionToken string, password []byte, newEmail string) error {
	return nil
}

func (self implAuth) ChangePassword(sessionToken string, oldPassword []byte, newPassword []byte) error {
	return nil
}

func createSignUpToken(domain, email string) (signUpToken string, err error) {
	return "", nil
}

func parseSignUpToken(signUpToken string) (domain, email string, err error) {
	return "", "", nil
}
