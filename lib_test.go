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

func TestCalcRSI(t *testing.T) {

	//testSlice := []float64{33.480000, 33.340000, 33.160000, 32.840000, 32.790001, 32.919998, 32.689999, 32.310001, 33.040001, 33.259998, 33.540001, 33.619999, 33.660000, 33.459999, 33.389999}
	testSlice := []float64{104.750000, 103.550003, 99.220001, 99.169998, 105.889999, 100.279999, 99.410004, 99.430000, 30.240000, 26.809999, 26.990000, 26.809999, 26.410000, 25.709999, 27.830000}
	RSI := CalcRSI(testSlice)
	t.Log("CalcRSI =", RSI)

	if RSI != 0.1248 {
		t.Errorf("Error: Testing CalcRSI")
	}
}

func TestHyperTanFunction(t *testing.T) {

	tanhResult := HyperTanFunction(0.123)
	t.Log("HyperTanFunction(0.123) = ", tanhResult)

	if tanhResult != 0.1248 {
		t.Errorf("Error: Testing HyperTanFunction")
	}
}

func TestCalcCreditSpreadCommissions(t *testing.T) {

	commission := CalcCreditSpreadCommissions(2)
	t.Log("CalcCreditSpreadCommissions( 2 ) = ", commission)

	if commission != 0.1248 {
		t.Errorf("Error: Testing CalcCreditSpreadCommissions")
	}
}

func TestCalcCreditSpreadMaxLoss(t *testing.T) {

  maxLoss := CalcCreditSpreadMaxLoss(5, 2, 480)
  t.Log("CalcCreditSpreadMaxLoss(5, 2, 480) = ", maxLoss)

  if maxLoss != 0.1248 {
    t.Errorf("Error: Testing CalcCreditSpreadMaxLoss")
  }
}
