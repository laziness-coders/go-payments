package adyen

import (
	"github.com/bos-hieu/go-payments/pkg/payments/types"
)

var (
	_ types.Client = (*adyen)(nil)
)

func NewClient() types.Client {
	result := &adyen{}
	return result
}

type adyen struct {
}

func (a adyen) Authorize(request *types.AuthorizeRequest) (response *types.AuthoriseResponse, err error) {
	panic("implement me")
}

func (a adyen) Capture(request *types.CaptureRequest) (response *types.CaptureResponse, err error) {
	panic("implement me")
}

func (a adyen) Void(request *types.VoidRequest) (response *types.VoidResponse, err error) {
	panic("implement me")
}

func (a adyen) Cancel(request *types.CancelRequest) (response *types.CancelResponse, err error) {
	panic("implement me")
}

func (a adyen) Refund(request *types.RefundRequest) (response *types.RefundResponse, err error) {
	panic("implement me")
}

func (a adyen) GenerateToken() string {
	panic("implement me")
}

