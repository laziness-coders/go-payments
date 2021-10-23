package checkoutdotcom

import (
	"github.com/bos-hieu/go-payments/pkg/payments/types"
)

var (
	_ types.Client = (*checkout)(nil)
)

func NewClient() types.Client {
	result := &checkout{}
	return result
}

type checkout struct {
}

func (c checkout) Authorize(request *types.AuthorizeRequest) (response *types.AuthoriseResponse, err error) {
	panic("implement me")
}

func (c checkout) Capture(request *types.CaptureRequest) (response *types.CaptureResponse, err error) {
	panic("implement me")
}

func (c checkout) Void(request *types.VoidRequest) (response *types.VoidResponse, err error) {
	panic("implement me")
}

func (c checkout) Cancel(request *types.CancelRequest) (response *types.CancelResponse, err error) {
	panic("implement me")
}

func (c checkout) Refund(request *types.RefundRequest) (response *types.RefundResponse, err error) {
	panic("implement me")
}

func (c checkout) GenerateToken() string {
	panic("implement me")
}
