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
	CREDIT_CARD PaymentMethod = iota
	BANK_TRANSFER
	ATM
)

type DeliveryMethod int

const (
	CTT_ECONOMIC DeliveryMethod = iota
	CTT_EXPRESS
)

type OrderStatus int

const (
	PAYMENT_PENDING OrderStatus = iota
	PAYMENT_UNAUTHORIZED
	PAYMENT_AUTHORIZED
	SENT
	DELIVERED
)
