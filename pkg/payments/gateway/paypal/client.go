package paypal

import (
	"github.com/bos-hieu/go-payments/pkg/payments/types"
)

var (
	_ types.Client = (*paypal)(nil)
)

func NewClient() types.Client {
	result := &paypal{}
	return result
}

type paypal struct {
}

func (p paypal) Authorize(request *types.AuthorizeRequest) (response *types.AuthoriseResponse, err error) {
	panic("implement me")
}

func (p paypal) Capture(request *types.CaptureRequest) (response *types.CaptureResponse, err error) {
	panic("implement me")
}

func (p paypal) Void(request *types.VoidRequest) (response *types.VoidResponse, err error) {
	panic("implement me")
}

func (p paypal) Cancel(request *types.CancelRequest) (response *types.CancelResponse, err error) {
	panic("implement me")
}

func (p paypal) Refund(request *types.RefundRequest) (response *types.RefundResponse, err error) {
	panic("implement me")
}

func (p paypal) GenerateToken() string {
	panic("implement me")
}
