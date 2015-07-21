package moneytransfer

type store interface {
	getAccount(id string) (Account, error)
	getAccountBalance(fromAccountId string) (int, error)
	createMoneyTransfer(mt MoneyTransfer) error
}
