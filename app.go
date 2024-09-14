package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	var expectedReturnRate float64 = 5.5
	var years float64 = 10

	var futureValue float64 = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Println("FutureValue is: ", futureValue)
}
