package affirm

import "github.com/bos-hieu/go-payments/pkg/payments"

var (
	_ payments.Client = (*affirm)(nil)
)

func NewClient() payments.Client {
	result := &affirm{}
	return result
}


type affirm struct {
}

func (a affirm) Authorize(request *payments.AuthorizeRequest) (response *payments.AuthoriseResponse, err error) {
	panic("implement me")
}

func (a affirm) Capture(request *payments.CaptureRequest) (response *payments.CaptureResponse, err error) {
	panic("implement me")
}

func (a affirm) Void(request *payments.VoidRequest) (response *payments.VoidResponse, err error) {
	panic("implement me")
}

func (a affirm) Cancel(request *payments.CancelRequest) (response *payments.CancelResponse, err error) {
	panic("implement me")
}

func (a affirm) Refund(request *payments.RefundRequest) (response *payments.RefundResponse, err error) {
	panic("implement me")
}

func (a affirm) GenerateToken() string {
	panic("implement me")
}

