package profitCalculator

import "fmt"

func InitializeProfitCalculator() {
	var revenue float64
	var expense float64
	var taxRate float64

	outputText("Enter Revenue: ")
	fmt.Scan(&revenue)

	outputText("Enter Expense: ")
	fmt.Scan(&expense)

	outputText("Enter Tax Rate: ")
	fmt.Scan(&taxRate)
  earningBeforeTax, profit, ratio := calculateEarningProfitAndRatio(revenue, expense, taxRate)

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
func calculateEarningProfitAndRatio(revenue, expense, taxRate float64)(float64, float64, float64){
	earningBeforeTax := revenue - expense
	profit := earningBeforeTax * (1 - taxRate/100)
	ratio := earningBeforeTax / profit
  return earningBeforeTax, profit, ratio
}
