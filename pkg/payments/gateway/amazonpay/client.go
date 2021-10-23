package amazonpay

import "github.com/bos-hieu/go-payments/pkg/payments"

var (
	_ payments.Client = (*amazonPay)(nil)
)

func NewClient() payments.Client {
	result := &amazonPay{}
	return result
}


type amazonPay struct {
}

func (a amazonPay) Authorize(request *payments.AuthorizeRequest) (response *payments.AuthoriseResponse, err error) {
	panic("implement me")
}

func (a amazonPay) Capture(request *payments.CaptureRequest) (response *payments.CaptureResponse, err error) {
	panic("implement me")
}

func (a amazonPay) Void(request *payments.VoidRequest) (response *payments.VoidResponse, err error) {
	panic("implement me")
}

func (a amazonPay) Cancel(request *payments.CancelRequest) (response *payments.CancelResponse, err error) {
	panic("implement me")
}

func (a amazonPay) Refund(request *payments.RefundRequest) (response *payments.RefundResponse, err error) {
	panic("implement me")
}

func (a amazonPay) GenerateToken() string {
	panic("implement me")
}

