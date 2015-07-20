package moneytransfer

import "fmt"

type MoneyTransfer interface {
	transfer(originAccountId, destinationAccountId string, amount int) error
}

type impl struct {
	store      store
	adminToken string
}

func New(store store, adminToken string) MoneyTransfer {
	return &impl{
		store:      store,
		adminToken: adminToken,
	}
}

func (self impl) transfer(originAccountId, destinationAccountId string, amount int) error {
	origin, err := self.store.getAccount(originAccountId)
	if err != nil {
		return err
	}

	// validate balance
	if origin.balance < amount {
		return fmt.Errorf("Error: Insufficient funds!")
	}

	destination, err := self.store.getAccount(destinationAccountId)
	if err != nil {
		return err
	}

	origin.balance -= amount
	destination.balance += amount

	// problem:
	// if between the "validate balance" comment above and the update(...) call
	// the origin account's balance changes to less than amount
	// this method should fail, but it does not
	return self.store.update([]Account{origin, destination})
}
