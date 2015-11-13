package finlib

import "math"

/**
  This calculates your monthly mortgage payment. Simple amortization formula.
  inputs:
  - loan amount
  - number of months it will take to pay off the loan
  - interest rate
  outputs:
  - monthly mortgage payment
*/
func CalcMortgagePayment(loanAmt float64, numOfMonths float64, interestRate float64) float64 {

	interestRatePerMonth := (interestRate / 100) / 12

	q := math.Pow(1+interestRatePerMonth, numOfMonths)

	return loanAmt * interestRatePerMonth * q / (q - 1)
}

/**
  This calculates the final amount your money grows to as interest compounds over time.
  Does not take into account inflation or consistent interest rate.
  inputs:
  - intial amount we start compounding from
  - interest rate (assumed, as we cannot predict what interest rates will be from year to year)
  - how many years will pass to let compounding work
  - how many times per year will account be compounded. Normally daily compounding, so 365
*/
func CalcCompountInterestAmount(initAmount float64, interestRate float64, durationInYears float64, numOfTimesCompoundedPerYear float64) float64 {

	return initAmount * math.Pow(1+(interestRate/100)/numOfTimesCompoundedPerYear, numOfTimesCompoundedPerYear*durationInYears)
}

/**
  Pretend you want a certain amount of money in the future. Calculate how much you need to start with today.
  Let's say I want to have $4,000,000 in 30 years with the money growing at 25% a year.
  How much initial money do I need to presently have for the scenario above to happen?
  inputs:
  - the money you have in the future
  - the interest rate you assume to be getting
  - the number of years it will take you to get to the future
*/
func CalcPresentValue(futureValue float64, annualInterestRate float64, durationInYears float64) float64 {

	return futureValue / math.Pow(1+annualInterestRate/100, durationInYears)
}
