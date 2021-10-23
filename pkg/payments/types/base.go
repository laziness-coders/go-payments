package types

type Address struct {
	City              string `json:"city"`
	Country           string `json:"country"`
	HouseNumberOrName string `json:"houseNumberOrName"`
	PostalCode        string `json:"postalCode"`
	StateOrProvince   string `json:"stateOrProvince"`
	Street            string `json:"street"`
}

type BillingAddress Address
type ShippingAddress Address

type Amount struct {
	Value    int    `json:"value"`
	Currency string `json:"currency"`
}

type CreditCard struct {

}

type AdditionalData struct {

}
