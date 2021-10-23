package stripe

import "github.com/bos-hieu/go-payments/pkg/payments/types"

var (
	_ types.Client = (*stripe)(nil)
)

func NewClient() types.Client {
	result := &stripe{}
	return result
}

type stripe struct {
}

func (s stripe) Authorize(request *types.AuthorizeRequest) (response *types.AuthoriseResponse, err error) {
	panic("implement me")
}

func (s stripe) Capture(request *types.CaptureRequest) (response *types.CaptureResponse, err error) {
	panic("implement me")
}

func (s stripe) Void(request *types.VoidRequest) (response *types.VoidResponse, err error) {
	panic("implement me")
}

func (s stripe) Cancel(request *types.CancelRequest) (response *types.CancelResponse, err error) {
	panic("implement me")
}

func (s stripe) Refund(request *types.RefundRequest) (response *types.RefundResponse, err error) {
	panic("implement me")
}

func (s stripe) GenerateToken() string {
	panic("implement me")
}
