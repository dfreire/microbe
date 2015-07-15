package auth

type User struct {
	email    string
	password string
}

type UnconfirmedUser User
