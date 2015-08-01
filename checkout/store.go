package checkout

type store interface {
	createOrder(order Order) error
	updateOrderStatus(order Order, status OrderStatus) error
}
