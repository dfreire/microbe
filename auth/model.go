package auth

type User struct {
	app      string
	email    string
	password string
}

type UnconfirmedUser User
