package strategies

import (
	"fmt"
	"github.com/bos-hieu/go-payments/pkg/payments/types"
)

type CreditCard struct {
}

func (c CreditCard) Authorize(request *types.AuthorizeRequest) (response *types.AuthoriseResponse, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CreditCard) Capture(request *types.CaptureRequest) (response *types.CaptureResponse, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CreditCard) Void(request *types.VoidRequest) (response *types.VoidResponse, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CreditCard) Cancel(request *types.CancelRequest) (response *types.CancelResponse, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CreditCard) Refund(request *types.RefundRequest) (response *types.RefundResponse, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CreditCard) GenerateToken() string {
	// TODO implement me
	panic("implement me")
}

func (c CreditCard) Pay() error {
	// Here you'll add the logic to pay with credit card

	fmt.Println("Processing purchase with CreditCard...")
	return nil
}

func NewCreditCard() *CreditCard {
	return &CreditCard{}
}
