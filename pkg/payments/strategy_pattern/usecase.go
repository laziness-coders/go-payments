package strategy_pattern

import (
	"errors"
	"fmt"
)

// Purchase is our goal/context which maintains a reference to all the
// possible strategies that the client can use to pay
type Purchase struct {
	paymentMethodStrategies map[int]PaymentMethodStrategy
}

func NewPurchase() Purchase {
	return Purchase{
		paymentMethodStrategies: make(map[int]PaymentMethodStrategy),
	}
}

// RegisterStrategy will let us register all the possible strategies that client can use
func (p *Purchase) RegisterStrategy(paymentMethod int, strategy PaymentMethodStrategy) {
	p.paymentMethodStrategies[paymentMethod] = strategy
}

// Process processes the purchase by doing the necessary validations
// and calling the Pay() method of selected strategy by the client
func (p Purchase) Process(paymentMethod int) error {
	// you can add logic to query and validate the order

	// in runtime, we chose the payment method strategy
	paymentMethodStrategy, ok := p.paymentMethodStrategies[paymentMethod]
	if !ok {
		return errors.New(fmt.Sprintf("Payment method %d is not supported", paymentMethod))
	}

	// after we got the strategy implementation we just use it by calling the Pay() method
	if _, err := paymentMethodStrategy.Authorize(nil); err != nil {
		return fmt.Errorf("purchase.Process.paymentMethodStrategy.Pay() failed: %v", err)
	}

	// here you may want to create an invoice, send a notification, etc

	fmt.Printf("Purchase with %d completed\n", paymentMethod)
	return nil
}
