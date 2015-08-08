package checkout

type Service interface {
	Checkout(order Order) error
}

type implService struct {
	store store
}

func New(store store) Service {
	return &implService{
		store: store,
	}
}

func (self implService) Checkout(order Order) error {
	order.Status = PAYMENT_PENDING
	return self.store.createOrder(order)
}

func (self implService) onPaymentAuthorized(orderId interface{}) error {
	order, err := self.store.getOrder(orderId)
	if err != nil {
		return err
	}

	return self.store.updateOrderStatus(order, PAYMENT_AUTHORIZED)
}

func (self implService) onPaymentUnauthorized(orderId interface{}) error {
	order, err := self.store.getOrder(orderId)
	if err != nil {
		return err
	}

	return self.store.updateOrderStatus(order, PAYMENT_UNAUTHORIZED)
}
