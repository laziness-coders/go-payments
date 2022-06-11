package strategy_pattern

import "github.com/bos-hieu/go-payments/pkg/payments/types"

// PaymentMethodStrategy will let us interchange the payment method in the context
// 1. Define the Contract Interface
type PaymentMethodStrategy interface {
	Pay() error
	Authorize(request *types.AuthorizeRequest) (response *types.AuthoriseResponse, err error)
	Capture(request *types.CaptureRequest) (response *types.CaptureResponse, err error)
	Void(request *types.VoidRequest) (response *types.VoidResponse, err error)
	Cancel(request *types.CancelRequest) (response *types.CancelResponse, err error)
	Refund(request *types.RefundRequest) (response *types.RefundResponse, err error)
	GenerateToken() string
}
