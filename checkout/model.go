package checkout

type Product struct {
	Id    interface{}
	Price int
}

type Address struct {
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

type Customer struct {
	UserId interface{}
	Name   string
}
