package checkout

type Checkout interface {
}

type implCheckout struct {
	store store
}

func New(store store) Checkout {
	return &implCheckout{}
}
