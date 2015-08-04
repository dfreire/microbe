package checkout

type store interface {
	getOrder(orderId interface{}) (Order, error)
	createOrder(order Order) error
	updatePaymentStatus(order Order, status PaymentStatus) error
	updateOrderStatus(order Order, status OrderStatus) error
}
