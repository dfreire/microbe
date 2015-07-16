package auth

import "time"

type User struct {
	Id          interface{}
	CreatedAt   time.Time
	Domain      string
	Email       string
	password    []byte
	isConfirmed bool
}

type Session struct {
	UserId    interface{}
	CreatedAt time.Time
}
