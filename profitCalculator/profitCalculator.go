package profitCalculator

import (
	"errors"
	"fmt"
	"os"
)
const storeResultFileName = "results.txt"

func storeResultsIntoFile(ebt, profit, ratio float64) error {
  output := fmt.Sprintf("earningBeforeTax: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
  writeFileErr := os.WriteFile(storeResultFileName, []byte(output), 0644)

  if writeFileErr != nil {
    return errors.New("Failed to store results into the file")
  }
  return nil
}

func InitializeProfitCalculator() {
	outputText("Enter Revenue: ")
	revenue, revenueErr := getUserInput()
	if revenueErr != nil {
		fmt.Println(revenueErr)
		return
	}

	outputText("Enter Expense: ")
	expense, expenseErr := getUserInput()
	if expenseErr != nil {
		fmt.Println(expenseErr)
		return
	}
  
	outputText("Enter Tax Rate: ")
	taxRate, taxRateErr := getUserInput()

	if taxRateErr != nil {
		fmt.Println(taxRateErr)
		return
	}

	earningBeforeTax, profit, ratio := calculateEarningProfitAndRatio(revenue, expense, taxRate)
  storeResultErr := storeResultsIntoFile(earningBeforeTax, profit, ratio)
  if storeResultErr != nil {
    fmt.Println(storeResultErr)
    return
  }

	//fmt.Println("Earning Before Tax is: ", earningBeforeTax)
	//fmt.Println("Profit After Tax is: ", profit)
	//fmt.Println("Ratio is: ", ratio)
	// fmt.Printf("Earnings before Tax is: %.2f\nProfit after Tax is: %.2f\nRatio is: %.2f", earningBeforeTax, profit, ratio)
	earningBeforeTaxFV := fmt.Sprintf("Earnings before Tax is: %.2f\n", earningBeforeTax)
	profitFV := fmt.Sprintf("Profit after Tax is: %.2f\n", profit)
	ratioFV := fmt.Sprintf("Ratio is: %.2f", ratio)
	print(earningBeforeTaxFV, profitFV, ratioFV)
	/*print(`This is the line break statement\n
	  Second Line`)*/

}

// custom function
func outputText(text string) {
	print(text)
}
func getUserInput() (float64, error) {
	var userInput float64
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return userInput, errors.New("Value must be a positive number and should not be zero")
	}
	return userInput, nil
}
func calculateEarningProfitAndRatio(revenue, expense, taxRate float64) (earningBeforeTax float64, profit float64, ratio float64) {
	earningBeforeTax = revenue - expense
	profit = earningBeforeTax * (1 - taxRate/100)
	ratio = earningBeforeTax / profit
	//return earningBeforeTax, profit, ratio
	return
}
