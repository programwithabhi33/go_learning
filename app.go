package main

import (
	"fmt"
	"math"
)

func main() {
	/* Normal way to declare variables
	var investmentAmount float64 = 1000
	var expectedReturnRate float64 = 5.5
	var years float64 = 10 */

	/* Shorter syntax for declaring variables
	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5 */

	/* Most preferred syntax for declaring variables */
	investmentAmount := 1000.0
	expectedReturnRate := 5.5
	years := 10.0

	var futureValue float64 = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Println("FutureValue is: ", futureValue)
}
