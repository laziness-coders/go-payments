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

func (a applePay) Authorize(request *types.AuthorizeRequest) (response *types.AuthoriseResponse, err error) {
	panic("implement me")
}

func (a applePay) Capture(request *types.CaptureRequest) (response *types.CaptureResponse, err error) {
	panic("implement me")
}

func (a applePay) Void(request *types.VoidRequest) (response *types.VoidResponse, err error) {
	panic("implement me")
}

func (a applePay) Cancel(request *types.CancelRequest) (response *types.CancelResponse, err error) {
	panic("implement me")
}

func (a applePay) Refund(request *types.RefundRequest) (response *types.RefundResponse, err error) {
	panic("implement me")
}

func (a applePay) GenerateToken() string {
	panic("implement me")
}
