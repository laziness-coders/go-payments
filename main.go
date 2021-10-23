package gopayments

import (
	"github.com/bos-hieu/go-payments/pkg/payments"
	"github.com/bos-hieu/go-payments/pkg/payments/constants/paymentgateway"
	"github.com/bos-hieu/go-payments/pkg/payments/gateway/adyen"
	"github.com/bos-hieu/go-payments/pkg/payments/gateway/amazonpay"
	"github.com/bos-hieu/go-payments/pkg/payments/gateway/paypal"
)

func NewClient(gateway int) payments.Client {
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
	_, err := client.Authorize(&payments.AuthorizeRequest{})
	if err != nil {

	}
}
