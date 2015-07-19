package moneytransfer

type store interface {
	getAccount(id string) (Account, error)
	update(accounts []Account) error
}
