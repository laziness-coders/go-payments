package authorizedotnet

import (
	"github.com/bos-hieu/go-payments/pkg/payments/types"
)

var (
	_ types.Client = (*authorizeDotNet)(nil)
)

func NewClient() types.Client {
	result := &authorizeDotNet{}
	return result
}


type authorizeDotNet struct {
}

func (a authorizeDotNet) Authorize(request *types.AuthorizeRequest) (response *types.AuthoriseResponse, err error) {
	panic("implement me")
}

func (a authorizeDotNet) Capture(request *types.CaptureRequest) (response *types.CaptureResponse, err error) {
	panic("implement me")
}

func (a authorizeDotNet) Void(request *types.VoidRequest) (response *types.VoidResponse, err error) {
	panic("implement me")
}

func (a authorizeDotNet) Cancel(request *types.CancelRequest) (response *types.CancelResponse, err error) {
	panic("implement me")
}

func (a authorizeDotNet) Refund(request *types.RefundRequest) (response *types.RefundResponse, err error) {
	panic("implement me")
}

func (a authorizeDotNet) GenerateToken() string {
	panic("implement me")
}

