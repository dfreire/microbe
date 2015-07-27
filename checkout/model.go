package checkout

type Order struct {
	Id                interface{}
	productIds        []interface{}
	customerId        interface{}
	invoiceAddressId  interface{}
	deliveryAddressId interface{}
	deliveryMethod    DeliveryMethod
	paymentMethod     PaymentMethod
	status            OrderStatus
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
