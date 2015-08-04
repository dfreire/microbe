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
	if err := self.store.createOrder(order); err != nil {
		return err
	}

	return self.store.updateOrderStatus(order, ORDER_STATUS_PENDING)
}

func (self implCheckout) onPaymentAuthorized(orderId interface{}) error {
	order, err := self.store.getOrder(orderId)
	if err != nil {
		return err
	}

	if err := self.store.updatePaymentStatus(order, PAYMENT_STATUS_AUTHORIZED); err != nil {
		return err
	}

	return self.store.updateOrderStatus(order, ORDER_STATUS_CONFIRMED)
}

func (self implCheckout) onPaymentUnauthorized(orderId interface{}) error {
	return nil
}
