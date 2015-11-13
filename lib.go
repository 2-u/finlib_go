package finlib

import "math"

func CalcMortgagePayment(loanAmt float64, numOfMonths float64, interestRate float64) float64 {

	interestRatePerMonth := (interestRate / 100) / 12

	q := math.Pow(1+interestRatePerMonth, numOfMonths)

	return loanAmt * interestRatePerMonth * q / (q - 1)
}
