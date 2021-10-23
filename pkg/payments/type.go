package payments

type Gateway interface {
	GetType() int
}

// PaymentRequest Is a contract to define the request is passed to Client Interface
type PaymentRequest struct {
}

// PaymentResponse is a contract to define the response from payment gateways
type PaymentResponse struct {
}

type AuthorizeRequest struct {
}

type AuthoriseResponse struct {
}

type CaptureRequest struct {
}

type CaptureResponse struct {
}

type VoidRequest struct {
}

type VoidResponse struct {
}

type CancelRequest struct {
}

type CancelResponse struct {
}

type RefundRequest struct {
}

type RefundResponse struct {
}

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

type Mandate struct {
	_                   struct{}
	Amount              string `json:"amount"`
	AmountRule          string `json:"amountRule"`
	BillingAttemptsRule string `json:"billingAttemptsRule"`
	BillingDay          string `json:"billingDay"`
	EndsAt              string `json:"endsAt"`
	Frequency           string `json:"frequency"`
	Remarks             string `json:"remarks"`
	StartsAt            string `json:"startsAt"`
}

type BankAccount struct {
	BankAccountNumber string `json:"bankAccountNumber"`
	BankCity          string `json:"bankCity"`
	BankLocationId    string `json:"bankLocationId"`
	BankName          string `json:"bankName,omitempty"`
	BIC               string `json:"bic,omitempty"`
	CountryCode       string `json:"countryCode"`
	IBAN              string `json:"iban,omitempty"`
	OwnerName         string `json:"ownerName,omitempty"`
	TaxId             string `json:"taxId,omitempty"`
}

// Client ...
type Client interface {
	Authorize(request *AuthorizeRequest) (response *AuthoriseResponse, err error)
	Capture(request *CaptureRequest) (response *CaptureResponse, err error)
	Void(request *VoidRequest) (response *VoidResponse, err error)
	Cancel(request *CancelRequest) (response *CancelResponse, err error)
	Refund(request *RefundRequest) (response *RefundResponse, err error)
	GenerateToken() string
}
