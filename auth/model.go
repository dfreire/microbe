package auth

type User struct {
	app      string
	email    string
	password []byte
}

type UnconfirmedUser User
