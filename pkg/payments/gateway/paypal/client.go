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

func (p paypal) Authorize(request *payments.AuthorizeRequest) (response *payments.AuthoriseResponse, err error) {
	panic("implement me")
}

func (p paypal) Capture(request *payments.CaptureRequest) (response *payments.CaptureResponse, err error) {
	panic("implement me")
}

func (p paypal) Void(request *payments.VoidRequest) (response *payments.VoidResponse, err error) {
	panic("implement me")
}

func (p paypal) Cancel(request *payments.CancelRequest) (response *payments.CancelResponse, err error) {
	panic("implement me")
}

func (p paypal) Refund(request *payments.RefundRequest) (response *payments.RefundResponse, err error) {
	panic("implement me")
}

func (p paypal) GenerateToken() string {
	panic("implement me")
}

