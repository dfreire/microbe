package auth

import "errors"

type Auth interface {
	signUp(app, email string, password []byte) (signUpToken string, err error)
	confirmSignUp(signUpToken string) error

	requestResetPasswordToken(app, email string) (resetPasswordToken string, err error)
	resetPassword(resetPasswordToken string, newPassword []byte) error

	signIn(app, email string, password []byte) (sessionToken string, err error)
	changeEmail(sessionToken string, password []byte, newEmail string) error
	changePassword(sessionToken string, oldPassword []byte, newPassword []byte) error
}

type impl struct {
	store store
}

func New(store store) Auth {
	return &impl{
		store: store,
	}
}

func (self impl) signUp(app, email string, password []byte) (signUpToken string, err error) {
	if err = self.assertCanCreateUser(app, email); err != nil {
		return "", err
	}

	signUpToken, err = createSignUpToken(app, email)
	if err != nil {
		return "", err
	}

	if err = self.store.createUnconfirmedUser(UnconfirmedUser{app, email, password}); err != nil {
		return "", err
	}

	return signUpToken, nil
}

func (self impl) confirmSignUp(signUpToken string) error {
	app, email, err := parseSignUpToken(signUpToken)
	if err != nil {
		return err
	}

	if err := self.assertCanCreateUser(app, email); err != nil {
		return err
	}

	unconfirmedUser, err := self.store.findUnconfirmedUser(app, email)
	if err != nil {
		return err
	}

	if err = self.store.createUser(User(unconfirmedUser)); err != nil {
		return err
	}

	err = self.store.removeUnconfirmedUser(unconfirmedUser)

	return nil
}

func (self impl) requestResetPasswordToken(app, email string) (resetPasswordToken string, err error) {
	return "", nil
}

func (self impl) resetPassword(resetPasswordToken string, newPassword []byte) error {
	return nil
}

func (self impl) signIn(app, email string, password []byte) (sessionToken string, err error) {
	return "", nil
}

func (self impl) changeEmail(sessionToken string, password []byte, newEmail string) error {
	return nil
}

func (self impl) changePassword(sessionToken string, oldPassword []byte, newPassword []byte) error {
	return nil
}

func (self impl) assertCanCreateUser(app, email string) error {
	exists, err := self.store.existsUser(app, email)
	if err != nil {
		return err
	} else if exists {
		return errors.New(ErrUserAlreadyExists)
	}

	return nil
}

func createSignUpToken(app, email string) (signUpToken string, err error) {
	return "", nil
}

func parseSignUpToken(signUpToken string) (app, email string, err error) {
	return "", "", nil
}
