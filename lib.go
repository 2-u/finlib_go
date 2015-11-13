package finlib

import "math"

func CalcMortgagePayment(loanAmt int, numOfMonths int, interestRate float64) float64 {

	interestRatePerMonth := (interestRate / 100) / 12

	q := math.Pow(1+interestRatePerMonth, float64(numOfMonths))

	return float64(loanAmt) * (interestRatePerMonth * q) / (q - 1)
}
