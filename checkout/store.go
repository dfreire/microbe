package checkout

type store interface {
	getOrder(orderId interface{}) (Order, error)
	createOrder(order Order) error
	updateOrderStatus(order Order, status OrderStatus) error
}

type implStore struct {
}
