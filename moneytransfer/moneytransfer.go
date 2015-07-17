package moneytransfer

type MoneyTransfer interface {
	transfer(origin Account, destination Account, amount int) error
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

func (self impl) transfer(origin Account, destination Account, amount int) error {
	return nil
}
