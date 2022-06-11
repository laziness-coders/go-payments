package strategy_pattern

import (
	"github.com/bos-hieu/go-payments/pkg/payments/constants/paymentgateway"
	"github.com/bos-hieu/go-payments/pkg/payments/strategy_pattern/strategies"
	"testing"
)

func TestNewPurchase(t *testing.T) {
	// client
	purchase := NewPurchase()
	purchase.RegisterStrategy(paymentgateway.Paypal, strategies.NewPayPal())
	purchase.RegisterStrategy(paymentgateway.Adyen, strategies.NewCreditCard())

	if err := purchase.Process(paymentgateway.Adyen); err != nil {
		t.Error(err)
	}
}
