package checkout

type store interface {
	getOrder(orderId interface{}) (Order, error)
	createOrder(order Order) error
	updateOrderStatus(order Order, status OrderStatus) error
}

type implStore struct {
}

func NewStore() store {
	return &implStore{}
}

func (self implStore) getOrder(orderId interface{}) (Order, error) {
	return Order{}, nil
}

func (self implStore) createOrder(order Order) error {
	return nil
}

func (self implStore) updateOrderStatus(order Order, status OrderStatus) error {
	return nil
}
