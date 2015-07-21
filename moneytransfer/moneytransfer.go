package moneytransfer

import "fmt"

type Service interface {
	transfer(originAccountId, destinationAccountId string, amount int) error
}

type impl struct {
	store      store
	adminToken string
}

func New(store store, adminToken string) Service {
	return &impl{
		store:      store,
		adminToken: adminToken,
	}
}

func (self impl) transfer(fromAccountId, toAccountId string, amount int) error {
	balance, err := self.store.getAccountBalance(fromAccountId)
	if err != nil {
		return err
	}

	// validate balance
	if balance < amount {
		return fmt.Errorf("Error: Insufficient funds!")
	}

	mt := MoneyTransfer{fromAccountId: fromAccountId, toAccountId: toAccountId, ammout: amount}

	// problem:
	// if between the "validate balance" comment above and the update(...) call
	// the origin account's balance changes to less than amount
	// this method should fail, but it does not
	return self.store.createMoneyTransfer(mt)
}
