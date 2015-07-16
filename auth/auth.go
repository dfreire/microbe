package auth

import (
	"errors"
	"time"
)

type Auth interface {
	signUp(domain, email string, password []byte) (signUpToken string, err error)
	confirmSignUp(signUpToken string) error

	requestResetPasswordToken(domain, email string) (resetPasswordToken string, err error)
	resetPassword(resetPasswordToken string, newPassword []byte) error

	signIn(domain, email string, password []byte) (sessionToken string, err error)
	changeEmail(sessionToken string, password []byte, newEmail string) error
	changePassword(sessionToken string, oldPassword []byte, newPassword []byte) error

	//getAllUsers(adminToken string) ([]User, error)
	//getUser(domain, email string) (User, error)
	//createUsers(adminToken, users []User) error
	//updateUsers(adminToken, users []User) error
	//removeUsers(adminToken, users []User) error
}

type impl struct {
	store      store
	adminToken string
}

func New(store store, adminToken string) Auth {
	return &impl{
		store:      store,
		adminToken: adminToken,
	}
}

func (self impl) signUp(domain, email string, password []byte) (signUpToken string, err error) {
	if err = self.assertCanCreateUser(domain, email); err != nil {
		return "", err
	}

	signUpToken, err = createSignUpToken(domain, email)
	if err != nil {
		return "", err
	}

	user := User{
		domain:      domain,
		email:       email,
		password:    password,
		createdAt:   time.Now(),
		isConfirmed: false,
	}
	if err = self.store.createUser(user); err != nil {
		return "", err
	}

	return signUpToken, nil
}

func (self impl) confirmSignUp(signUpToken string) error {
	domain, email, err := parseSignUpToken(signUpToken)
	if err != nil {
		return err
	}

	if err := self.assertCanCreateUser(domain, email); err != nil {
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

func (self impl) requestResetPasswordToken(domain, email string) (resetPasswordToken string, err error) {
	return "", nil
}

func (self impl) resetPassword(resetPasswordToken string, newPassword []byte) error {
	return nil
}

func (self impl) signIn(domain, email string, password []byte) (sessionToken string, err error) {
	return "", nil
}

func (self impl) changeEmail(sessionToken string, password []byte, newEmail string) error {
	return nil
}

func (self impl) changePassword(sessionToken string, oldPassword []byte, newPassword []byte) error {
	return nil
}

func (self impl) assertCanCreateUser(domain, email string) error {
	user, err := self.store.getUser(domain, email)
	if err != nil {
		return err
	} else if exists {
		return errors.New(ErrEquivalentEntityExists)
	}

	return nil
}

func createSignUpToken(domain, email string) (signUpToken string, err error) {
	return "", nil
}

func parseSignUpToken(signUpToken string) (domain, email string, err error) {
	return "", "", nil
}
