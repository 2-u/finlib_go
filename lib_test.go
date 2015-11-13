package finlib

import (
	"testing"
)

func TestCalcMortgagePayment(t *testing.T) {

	mortgagePayment := CalcMortgagePayment(240000, 360, 3.5)
	t.Log("mortgagePayment = ")
	t.Log(mortgagePayment)

	if mortgagePayment != 0.1248 {
		t.Errorf("Error: Testing CalcMortgagePayment")
	}
}

func TestCalcCompountInterestAmount(t *testing.T) {

  amoountWithCompoundedInterest := CalcCompountInterestAmount(12000, 5, 30, 365)
  t.Log("amoountWithCompoundedInterest = ")
  t.Log(amoountWithCompoundedInterest)

  // Force it to be wrong
  if amoountWithCompoundedInterest != 0.1248 {
    t.Errorf("Error: Testing CalcCompountInterestAmount")
  }
}

func TestCalcPresentValue(t *testing.T) {

  presentValue := CalcPresentValue( 4000000, 25, 15 )
  t.Log("presentValue = ")
  t.Log(presentValue)

  if presentValue != 0.1248 {
    t.Errorf("Error: Testing TestCalcPresentValue")
  }
}
