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

func CalcPercentChange(fromVal float64, toVal float64) float64 {

	return (toVal - fromVal) / fromVal
}

func CalcAvg(dataSlice []float64) float64 {

	sum := 0.0

	for i := range dataSlice {
		sum += float64(dataSlice[i])
	}

	return float64(sum) / float64(len(dataSlice))
}

func CalcRSI(dataSlice15 []float64) float64 {

	upSum := 0.0
	downSum := 0.0

	for i := 0; i < len(dataSlice15)-1; i++ {

		curPrice := dataSlice15[i]
		tomPrice := dataSlice15[i+1]

		if curPrice < tomPrice {
			upSum += tomPrice - curPrice
		} else {
			downSum += curPrice - tomPrice
		}
	}

	avgUp := upSum / 14.0
	avgDown := downSum / 14.0
	relativeStrength := avgUp / avgDown

	return 100 - 100/(1+relativeStrength)
}

func HyperTanFunction(x float64) float64 {

	// This function is used in neural network programming
	// as an activation function. Use this for neural nets
	// with range from -1 to 1

	if x < -20.0 {
		return -1.0 // approximation is correct to 30 decimals
	} else if x > 20.0 {
		return 1.0
	}

	return math.Tanh(x)
}

func CalcCreditSpreadCommissions(numOfSpreads float64) float64 {

	// TD Ameritrade costs $9.99 per trade + 0.75 per contract
	return numOfSpreads*2*0.75 + 9.99
}

func CalcCreditSpreadMaxLoss(strikesWidth float64, numOfSpreads float64, creditCollectedAmt float64) float64 {

	// This function calculates the maximum loss for
	// Bull Put and Bear Call spreads (includes commissions).
	return numOfSpreads*(strikesWidth-creditCollectedAmt)*100 + 2*CalcCreditSpreadCommissions(numOfSpreads)
}

func CalcCreditSpreadMaxProfit(numOfSpreads float64, creditCollectedAmt float64) float64 {

	// Calculates the max profit for
	// Bull Put and Bear Call spreads (includes commissions).
	return numOfSpreads*creditCollectedAmt*100 - CalcCreditSpreadCommissions(numOfSpreads)
}

func CalcCreditSpreadBreakEvenPrice(shortOptionStrikePrice float64, longOptionStrikePrice float64, creditCollectedAmt float64) float64 {

	if shortOptionStrikePrice < longOptionStrikePrice {
		return shortOptionStrikePrice + creditCollectedAmt
	}

	return shortOptionStrikePrice - creditCollectedAmt
}

func CalcCreditSpreadTradeExitPrice(soldPrice float64, boughtBackPrice float64, numOfSpreads float64) float64 {

	// Calculates the profit or loss made when exiting your
	// Bull Put or Bear Call Spread (includes commissions on trade entry and exit).
	return (soldPrice-boughtBackPrice)*100*numOfSpreads - 2*CalcCreditSpreadCommissions(numOfSpreads)
}

func CalcCreditSpreadProbabilityOfProfit(creditCollectedAmt float64, strikePriceWidth float64) float64 {

	return 100.0 - ((creditCollectedAmt / strikePriceWidth) * 100.0)
}
