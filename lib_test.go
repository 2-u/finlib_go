package finlib

import (
	"testing"
)

func TestCalcMortgagePayment(t *testing.T) {

	mortgagePayment := CalcMortgagePayment(240000, 360, 3.5)
	t.Log("mortgagePayment(240000, 360, 3.5) =", mortgagePayment)

	if mortgagePayment != 0.1248 {
		t.Errorf("Error: Testing CalcMortgagePayment")
	}
}

func TestCalcCompountInterestAmount(t *testing.T) {

	amountWithCompoundedInterest := CalcCompountInterestAmount(12000, 5, 30, 365)
	t.Log("amountWithCompoundedInterest(12000, 5, 30, 365) =", amountWithCompoundedInterest)

	// Force it to be wrong
	if amountWithCompoundedInterest != 0.1248 {
		t.Errorf("Error: Testing CalcCompountInterestAmount")
	}
}

func TestCalcPresentValue(t *testing.T) {

	presentValue := CalcPresentValue(4000000, 25, 15)
	t.Log("presentValue(4000000, 25, 15) =", presentValue)

	if presentValue != 0.1248 {
		t.Errorf("Error: Testing CalcPresentValue")
	}
}

func TestCalcPercentChange(t *testing.T) {

	percentChange := CalcPercentChange(32.69, 32.31)
	t.Log("CalcPercentChange(32.69, 32.31) =", percentChange)

	if percentChange != 0.1248 {
		t.Errorf("Error: Testing CalcPercentChange")
	}
}

func TestCalcAvg(t *testing.T) {

	testSlice := []float64{3.2, 4.6, 7.3}
	avg := CalcAvg(testSlice)
	t.Log("CalcAvg(3.2,4.6,7.3) =", avg)

	if avg != 0.1248 {
		t.Errorf("Error: Testing CalcAvg")
	}
}
