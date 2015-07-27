package checkout

import "time"

type Order struct {
	Id                interface{}
	CreatedAt         time.Time
	CustomerId        interface{}
	InvoiceAddressId  interface{}
	DeliveryAddressId interface{}
	DeliveryMethod    DeliveryMethod
	Lines             []OrderLine
	PaymentMethod     PaymentMethod
	Status            OrderStatus
}

type OrderLine struct {
	ProductId interface{}
	Quantity  int
	Price     int
}

type Product struct {
	Id    interface{}
	Price int
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
	CreditCard PaymentMethod = iota
	BankTransfer
	ATMTrasnfer
)

type DeliveryMethod int

const (
	CTTEconomic DeliveryMethod = iota
	CTTExpress
)

type OrderStatus int

const (
	PaymentPending OrderStatus = iota
	PaymentUnauthorized
	Confirmed
	Sent
	Delivered
)
