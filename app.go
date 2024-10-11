package main

/*import (
	"fmt"
	"math"
)*/
//import "first-app/profitCalculator"
import "first-app/basicFinancial"

func main() {
	//Normal way to declare variables
	/*var investmentAmount float64 = 1000
	var expectedReturnRate float64 = 5.5
	var years float64 = 10*/

	/* Shorter syntax for declaring variables
	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5 */

	/* Most preferred syntax for declaring variables */
  /*const inflationRate = 2.5
  var investmentAmount, years float64
	expectedReturnRate := 5.5

  fmt.Print("Enter InvestmentAmount: ")
  fmt.Scan(&investmentAmount)

  fmt.Print("Enter expectedReturnRate: ")
  fmt.Scan(&expectedReturnRate)

  fmt.Print("Enter no.of Years: ")
  fmt.Scan(&years)

	var futureValue float64 = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
  futureRealValue := futureValue/math.Pow(1+inflationRate/100, years)

	fmt.Println("FutureValue is: ", futureValue)
  fmt.Println("FutureRealValue is: ", futureRealValue)*/
  //profitCalculator.InitializeProfitCalculator()
  basicFinancial.ShowChoices()
}
