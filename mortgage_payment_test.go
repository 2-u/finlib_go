package finlib

import (
	"testing"
)

func TestCalcMortgagePayment(t *testing.T) {

	mortgagePayment := CalcMortgagePayment(100000, 360, 3.5)
	t.Log("mortgagePayment = ")
	t.Log(mortgagePayment)

	if mortgagePayment != 0.1248 {
		t.Errorf("Error: Testing CalcMortgagePayment")
	}
}
