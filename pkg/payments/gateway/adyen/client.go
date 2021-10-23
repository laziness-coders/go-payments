package adyen

import "github.com/bos-hieu/go-payments/pkg/payments"

var (
	_ payments.Client = (*adyen)(nil)
)

func NewClient() payments.Client {
	result := &adyen{}
	return result
}

type adyen struct {
}

func (a adyen) Authorize(request *payments.AuthorizeRequest) (response *payments.AuthoriseResponse, err error) {
	panic("implement me")
}

func (a adyen) Capture(request *payments.CaptureRequest) (response *payments.CaptureResponse, err error) {
	panic("implement me")
}

func (a adyen) Void(request *payments.VoidRequest) (response *payments.VoidResponse, err error) {
	panic("implement me")
}

func (a adyen) Cancel(request *payments.CancelRequest) (response *payments.CancelResponse, err error) {
	panic("implement me")
}

func (a adyen) Refund(request *payments.RefundRequest) (response *payments.RefundResponse, err error) {
	panic("implement me")
}

func (a adyen) GenerateToken() string {
	panic("implement me")
}

