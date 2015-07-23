package auth

import (
	"errors"
	"time"
)

type Auth interface {
	SignUp(domain, email string, password []byte) (signUpToken string, err error)
	ConfirmSignUp(signUpToken string) error

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

func (self impl) SignUp(domain, email string, password []byte) (signUpToken string, err error) {
	if err = self.assertCanCreateUser(domain, email); err != nil {
		return "", err
	}

	signUpToken, err = createSignUpToken(domain, email)
	if err != nil {
		return "", err
	}

	user := User{
		CreatedAt:   time.Now(),
		Domain:      domain,
		Email:       email,
		password:    password,
		isConfirmed: false,
	}
	if err = self.store.createUser(user); err != nil {
		return "", err
	}

	return signUpToken, nil
}

func (self impl) ConfirmSignUp(signUpToken string) error {
	domain, email, err := parseSignUpToken(signUpToken)
	if err != nil {
		return err
	}

	user, err := self.store.getUser(domain, email)
	if err != nil {
		return err
	}

	user.isConfirmed = true
	if err = self.store.updateUser(user, []string{"isConfirmed"}); err != nil {
		return err
	}

	return nil
}

func (self impl) RequestResetPasswordToken(domain, email string) (resetPasswordToken string, err error) {
	return "", nil
}

func (self impl) ResetPassword(resetPasswordToken string, newPassword []byte) error {
	return nil
}

func (self impl) SignIn(domain, email string, password []byte) (sessionToken string, err error) {
	return "", nil
}

func (self impl) ChangeEmail(sessionToken string, password []byte, newEmail string) error {
	return nil
}

func (self impl) ChangePassword(sessionToken string, oldPassword []byte, newPassword []byte) error {
	return nil
}

func (self impl) assertCanCreateUser(domain, email string) error {
	_, err := self.store.getUser(domain, email)
	if err != nil {
		if err.Error() == ErrEntityNotFound {
			return nil
		} else {
			return err
		}
	}

	return errors.New(ErrEntityAlreadyExists)
}

func createSignUpToken(domain, email string) (signUpToken string, err error) {
	return "", nil
}

func parseSignUpToken(signUpToken string) (domain, email string, err error) {
	return "", "", nil
}
