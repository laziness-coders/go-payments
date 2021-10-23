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

func (a authorizeDotNet) Authorize(request *payments.AuthorizeRequest) (response *payments.AuthoriseResponse, err error) {
	panic("implement me")
}

func (a authorizeDotNet) Capture(request *payments.CaptureRequest) (response *payments.CaptureResponse, err error) {
	panic("implement me")
}

func (a authorizeDotNet) Void(request *payments.VoidRequest) (response *payments.VoidResponse, err error) {
	panic("implement me")
}

func (a authorizeDotNet) Cancel(request *payments.CancelRequest) (response *payments.CancelResponse, err error) {
	panic("implement me")
}

func (a authorizeDotNet) Refund(request *payments.RefundRequest) (response *payments.RefundResponse, err error) {
	panic("implement me")
}

func (a authorizeDotNet) GenerateToken() string {
	panic("implement me")
}

