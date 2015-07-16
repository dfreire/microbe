package auth

import "time"

type User struct {
	id          interface{}
	domain      string
	email       string
	password    []byte
	createdAt   time.Time
	isConfirmed bool
}

type Session struct {
	domain    string
	userId    interface{}
	createdAt time.Time
}
