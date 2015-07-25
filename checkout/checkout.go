package checkout

type Checkout interface {
	PlaceOrder(order Order) error
}

type implCheckout struct {
	store store
}

func New(store store) Checkout {
	return &implCheckout{
		store: store,
	}
}

func (self implCheckout) PlaceOrder(order Order) error {
	return nil
}
