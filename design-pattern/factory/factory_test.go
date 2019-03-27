package factory

import (
	"strings"
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method of type 'cash' must exist!")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("the cash method message wasn't correct!")
	}

	t.Log("LOG:", msg)
}


func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatal("A payment method of type 'cash' must exist!")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("the cash method message wasn't correct!")
	}

	t.Log("LOG:", msg)
}