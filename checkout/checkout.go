package checkout

type Checkout interface {
	Checkout(customer Customer, address Address, products []Product) error
}

type implCheckout struct {
	store store
}

func New(store store) Checkout {
	return &implCheckout{
		store: store,
	}
}

func (self implCheckout) Checkout(customer Customer, address Address, products []Product) error {
	return nil
}
