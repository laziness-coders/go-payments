package payments

type Gateway interface {
	GetType() int
}

// PaymentRequest Is a contract to define the request is passed to Payment Interface
type PaymentRequest interface {
	GetRequest() interface{}
}

// PaymentResponse is a contract to define the response from payment gateways
type PaymentResponse interface {
	GetResponse() interface{}
}

// Payment ...
type Payment interface {
	Authorize(request PaymentRequest) (response PaymentResponse, err error)
	Capture(request PaymentRequest) (response PaymentResponse, err error)
	Void(request PaymentRequest) (response PaymentResponse, err error)
	Cancel(request PaymentRequest) (response PaymentResponse, err error)
	Refund(request PaymentRequest) (response PaymentResponse, err error)
	GenerateToken() string
}
