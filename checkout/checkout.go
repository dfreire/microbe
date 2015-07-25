package checkout

type Checkout interface {
	Checkout(products []Product) error
}

type implCheckout struct {
	store store
}

func New(store store) Checkout {
	return &implCheckout{
		store: store,
	}
}

func (self implCheckout) Checkout(products []Product) error {
	return nil
}
