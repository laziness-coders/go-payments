package gopayments

import (
	"github.com/bos-hieu/go-payments/pkg/payments/constants/paymentgateway"
	"github.com/bos-hieu/go-payments/pkg/payments/gateway/adyen"
	"github.com/bos-hieu/go-payments/pkg/payments/gateway/amazonpay"
	"github.com/bos-hieu/go-payments/pkg/payments/gateway/paypal"
	"github.com/bos-hieu/go-payments/pkg/payments/types"
)

func NewClient(gateway int) types.Client {
	switch gateway {
	case paymentgateway.Adyen:
		return adyen.NewClient()
	case paymentgateway.AmazonPay:
		return amazonpay.NewClient()
	case paymentgateway.Paypal:
		return paypal.NewClient()
	}

	return nil
}

func HandleClient() {
	client := NewClient(paymentgateway.Adyen)
	_, err := client.Authorize(&types.AuthorizeRequest{})
	if err != nil {

	}
}
