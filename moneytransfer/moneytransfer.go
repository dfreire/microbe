package moneytransfer

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

	destination, err := self.store.getAccount(destinationAccountId)
	if err != nil {
		return err
	}

	origin.balance -= amount
	destination.balance += amount
	return self.store.update([]Account{origin, destination})
}
