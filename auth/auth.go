package auth

import "errors"

type Auth interface {
	signUp(email string, password string) (signUpToken string, err error)
	confirmSignUp(signUpToken string) error
	signIn(email string, password string) (sessionToken string, err error)
	requestResetPasswordToken(email string) (resetPasswordToken string, err error)
	resetPassword(resetPasswordToken string) error
	changeEmail(sessionToken string, password string, newEmail string) error
	changePassword(sessionToken string, oldPassword string, newPassword string) error
}

type impl struct {
	store store
}

func New(store store) Auth {
	return &impl{
		store: store,
	}
}

func (self impl) signUp(email string, password string) (signUpToken string, err error) {
	if err = self.assertCanCreateUser(email); err != nil {
		return "", err
	}

	signUpToken, err = createSignUpToken(email)
	if err != nil {
		return "", err
	}

	if err = self.store.putPendingUser(email, password); err != nil {
		return "", err
	}

	return signUpToken, nil
}

func (self impl) confirmSignUp(signUpToken string) error {
	email, err := parseSignUpToken(signUpToken)
	if err != nil {
		return err
	}

	if err := self.assertCanCreateUser(email); err != nil {
		return err
	}

	pendingUser, err := self.store.getPendingUser(email)
	if err != nil {
		return err
	}

	if err = self.store.createUser(User(pendingUser)); err != nil {
		return err
	}

	err = self.store.delPendingUser(email)

	return nil
}

func (self impl) signIn(email string, password string) (sessionToken string, err error) {
	return "", nil
}

func (self impl) requestResetPasswordToken(email string) (resetPasswordToken string, err error) {
	return "", nil
}

func (self impl) resetPassword(resetPasswordToken string) error {
	return nil
}

func (self impl) changeEmail(sessionToken string, password string, newEmail string) error {
	return nil
}

func (self impl) changePassword(sessionToken string, oldPassword string, newPassword string) error {
	return nil
}

func (self impl) assertCanCreateUser(email string) error {
	hasUserWithEmail, err := self.store.hasUserWithEmail(email)
	if err != nil {
		return err
	} else if hasUserWithEmail {
		return errors.New(ErrEmailAlreadyUsed)
	}

	return nil
}

func createSignUpToken(email string) (signUpToken string, err error) {
	return "", nil
}

func parseSignUpToken(signUpToken string) (email string, err error) {
	return "", nil
}
