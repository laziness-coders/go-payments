package applepay

import (
	"github.com/bos-hieu/go-payments/pkg/payments/types"
)

var (
	_ types.Client = (*applePay)(nil)
)

func NewClient() types.Client {
	result := &applePay{}
	return result
}

type applePay struct {
}

func (a applePay) Authorize(request *payments.AuthorizeRequest) (response *payments.AuthoriseResponse, err error) {
	panic("implement me")
}

func (a applePay) Capture(request *payments.CaptureRequest) (response *payments.CaptureResponse, err error) {
	panic("implement me")
}

func (a applePay) Void(request *payments.VoidRequest) (response *payments.VoidResponse, err error) {
	panic("implement me")
}

func (a applePay) Cancel(request *payments.CancelRequest) (response *payments.CancelResponse, err error) {
	panic("implement me")
}

func (a applePay) Refund(request *payments.RefundRequest) (response *payments.RefundResponse, err error) {
	panic("implement me")
}

func (a applePay) GenerateToken() string {
	panic("implement me")
}

