package checkout

type Order struct {
	Id                interface{}
	productIds        []interface{}
	customerId        interface{}
	invoiceAddressId  interface{}
	deliveryAddressId interface{}
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

type OrderStatus int

const (
	Placed OrderStatus = iota
	PaymentPending
	PaymentUnauthorized
	Delivered
)
