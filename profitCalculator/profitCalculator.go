package profitCalculator

import "fmt"

func InitializeProfitCalculator() {
  var revenue float64
  var expense float64
  var taxRate float64

  print("Enter Revenue: ")
  fmt.Scan(&revenue)

  print("Enter Expense: ")
  fmt.Scan(&expense)

  print("Enter Tax Rate: ")
  fmt.Scan(&taxRate)
  
  earningBeforeTax := revenue - expense
  profit := earningBeforeTax * (1-taxRate/100)
  ratio := earningBeforeTax/profit

  fmt.Println("Earning Before Tax is: ", earningBeforeTax)
  fmt.Println("Profit After Tax is: ", profit)
  fmt.Println("Ratio is: ", ratio)
}
