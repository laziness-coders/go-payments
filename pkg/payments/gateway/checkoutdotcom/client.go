package checkoutdotcom

import "github.com/bos-hieu/go-payments/pkg/payments"

var (
	_ payments.Client = (*checkout)(nil)
)

func NewClient() payments.Client {
	result := &checkout{}
	return result
}

type checkout struct {
}

func (c checkout) Authorize(request *payments.AuthorizeRequest) (response *payments.AuthoriseResponse, err error) {
	panic("implement me")
}

func (c checkout) Capture(request *payments.CaptureRequest) (response *payments.CaptureResponse, err error) {
	panic("implement me")
}

func (c checkout) Void(request *payments.VoidRequest) (response *payments.VoidResponse, err error) {
	panic("implement me")
}

func (c checkout) Cancel(request *payments.CancelRequest) (response *payments.CancelResponse, err error) {
	panic("implement me")
}

func (c checkout) Refund(request *payments.RefundRequest) (response *payments.RefundResponse, err error) {
	panic("implement me")
}

func (c checkout) GenerateToken() string {
	panic("implement me")
}

