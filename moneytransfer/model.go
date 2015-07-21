package moneytransfer

type Account struct {
	Id string
}

type MoneyTransfer struct {
	fromAccountId string
	toAccountId   string
	ammout        int
}
