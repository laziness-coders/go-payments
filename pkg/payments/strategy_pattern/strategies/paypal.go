package strategies

import (
	"fmt"
	"github.com/bos-hieu/go-payments/pkg/payments/types"
)

// PayPal is an implementation of the PaymentMethodStrategy
// interfaith necessary logic to pay with PayPal
type PayPal struct {
}

func (p PayPal) Authorize(request *types.AuthorizeRequest) (response *types.AuthoriseResponse, err error) {
	// TODO implement me
	panic("implement me")
}

func (p PayPal) Capture(request *types.CaptureRequest) (response *types.CaptureResponse, err error) {
	// TODO implement me
	panic("implement me")
}

func (p PayPal) Void(request *types.VoidRequest) (response *types.VoidResponse, err error) {
	// TODO implement me
	panic("implement me")
}

func (p PayPal) Cancel(request *types.CancelRequest) (response *types.CancelResponse, err error) {
	// TODO implement me
	panic("implement me")
}

func (p PayPal) Refund(request *types.RefundRequest) (response *types.RefundResponse, err error) {
	// TODO implement me
	panic("implement me")
}

func (p PayPal) GenerateToken() string {
	// TODO implement me
	panic("implement me")
}

func (p PayPal) Pay() error {
	// Here you'll add the logic to pay with PayPal

	fmt.Println("Processing purchase with PayPal...")

	return nil
}

func NewPayPal() *PayPal {
	return &PayPal{}
}
