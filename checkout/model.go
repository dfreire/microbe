package checkout

import "time"

type Order struct {
	Id                interface{}
	CreatedAt         time.Time
	CustomerId        interface{}
	InvoiceAddressId  interface{}
	DeliveryAddressId interface{}
	DeliveryMethod    DeliveryMethod
	PaymentMethod     PaymentMethod
	Lines             []OrderLine
	Status            OrderStatus
}

type OrderLine struct {
	ProductId interface{}
	Quantity  int
	Price     int
}

type Product struct {
	Id                interface{}
	Price             int
	StockQuantity     int
	AllocatedQuantity int
}

func (self Product) getAvailableQuantity() int {
	return self.StockQuantity - self.AllocatedQuantity
}

type Customer struct {
	UserId    interface{}
	Name      string
	AdressIds []interface{}
}

type Address struct {
	Id           interface{}
	FullName     string
	CompanyName  string
	AddressLine1 string
	AddressLine2 string
	TownOrCity   string
	County       string
	Postcode     string
	Country      string
	PhoneNumber  string
}

type PaymentMethod int

const (
	PAYMENT_METHOD_CREDIT_CARD PaymentMethod = iota
	PAYMENT_METHOD_BANK_TRANSFER
	PAYMENT_METHOD_ATM
)

type DeliveryMethod int

const (
	DELIVERY_METHOD_CTT_ECONOMIC DeliveryMethod = iota
	DELIVERY_METHOD_CTT_EXPRESS
)

type OrderStatus int

const (
	ORDER_STATUS_PAYMENT_PENDING OrderStatus = iota
	ORDER_STATUS_PAYMENT_UNAUTHORIZED
	ORDER_STATUS_PAYMENT_AUTHORIZED
	ORDER_STATUS_SENT
	ORDER_STATUS_DELIVERED
)
