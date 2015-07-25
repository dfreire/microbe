package checkout

type Checkout interface {
}

type implCheckout struct {
}

func New() Checkout {
	return &implCheckout{}
}
