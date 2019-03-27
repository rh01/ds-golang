package factory

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float64) string
}

const (
	Cash      = 1
	DebitCard = 2
)

type CashPM struct {
}

func (c *CashPM) Pay(amount float64) string {
	return fmt.Sprint("paid using cash")
}

type DebitCardPM struct {
}

func (d *DebitCardPM) Pay(amount float64) string {
	return fmt.Sprintf("paid using debit card")
}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recongnized\n", m))
	}
}
