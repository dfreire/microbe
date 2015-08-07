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
	order.Status = PAYMENT_PENDING
	return self.store.createOrder(order)
}

func (self implCheckout) onPaymentAuthorized(orderId interface{}) error {
	order, err := self.store.getOrder(orderId)
	if err != nil {
		return err
	}

	return self.store.updateOrderStatus(order, PAYMENT_AUTHORIZED)
}

func (self implCheckout) onPaymentUnauthorized(orderId interface{}) error {
	order, err := self.store.getOrder(orderId)
	if err != nil {
		return err
	}

	return self.store.updateOrderStatus(order, PAYMENT_UNAUTHORIZED)
}
