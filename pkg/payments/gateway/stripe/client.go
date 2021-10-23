package stripe

import "github.com/bos-hieu/go-payments/pkg/payments"

var (
	_ payments.Client = (*stripe)(nil)
)

func NewClient() payments.Client {
	result := &stripe{}
	return result
}

type stripe struct {
}

func (s stripe) Authorize(request *payments.AuthorizeRequest) (response *payments.AuthoriseResponse, err error) {
	panic("implement me")
}

func (s stripe) Capture(request *payments.CaptureRequest) (response *payments.CaptureResponse, err error) {
	panic("implement me")
}

func (s stripe) Void(request *payments.VoidRequest) (response *payments.VoidResponse, err error) {
	panic("implement me")
}

func (s stripe) Cancel(request *payments.CancelRequest) (response *payments.CancelResponse, err error) {
	panic("implement me")
}

func (s stripe) Refund(request *payments.RefundRequest) (response *payments.RefundResponse, err error) {
	panic("implement me")
}

func (s stripe) GenerateToken() string {
	panic("implement me")
}
